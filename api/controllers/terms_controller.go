package controllers

import (
	"dictionary/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// FindTerms - Find all terms
// GET /terms
func FindTerms(c *gin.Context) {
	var terms []models.Term
	models.DB.Preload("RelatedTerms").Find(&terms)

	c.JSON(http.StatusOK, gin.H{"data": terms})
}

// FindTerms - Find term by id
// GET /terms/:term_id
func FindTerm(c *gin.Context) {
	var ok bool
	termID := c.Param("term_id")
	if termID != "" {
		id := uuid.FromStringOrNil(termID)
		if id != uuid.Nil {
			var term models.Term
			err := models.DB.Preload("RelatedTerms").Model(&models.Term{}).Where("id = ?", id).Take(&term).Error
			if err == nil {
				c.JSON(http.StatusOK, gin.H{"data": term})
				ok = true
			}
		}
	}
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"data": "id not found"})
	}
}

// TermInput - DTO
type TermInput struct {
	Name         string               `json:"name" binding:"required"`
	Pronuciation string               `json:"pronuciation"`
	Definition   string               `json:"definition"`
	Synonyms     string               `json:"synonyms"`
	RelatedTerms []models.RelatedTerm `json:"related_terms"`
}

// CreateTerm - Create new term
// POST /terms
func CreateTerm(c *gin.Context) {
	input, err := bindTerm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := GetLoggedUserID(c)

	// Create term
	term := models.Term{Name: input.Name, Pronuciation: input.Pronuciation,
		Definition: input.Pronuciation, Synonyms: input.Synonyms,
		RelatedTerms: input.RelatedTerms,
		CreatedByID:  userID}
	err = models.DB.Create(&term).Error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"data": term})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": term})
	}
}

// UpdateTerm - Update term
// PUT /terms/:term_id
func UpdateTerm(c *gin.Context) {
	termID := c.Param("term_id")
	if termID != "" {
		id := uuid.FromStringOrNil(termID)
		if id != uuid.Nil && existsTerm(id) {
			input, err := bindTerm(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			userID := GetLoggedUserID(c)

			// Create term
			term := models.Term{ID: id, Name: input.Name, Pronuciation: input.Pronuciation,
				Definition: input.Pronuciation, Synonyms: input.Synonyms,
				RelatedTerms: input.RelatedTerms,
				UpdatedByID:  userID}
			err = models.DB.Save(&term).Error
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{"data": term})
			} else {
				c.JSON(http.StatusOK, gin.H{"data": term})
			}
		} else {
			c.JSON(http.StatusNotFound, gin.H{"data": "id not found"})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "term id is required"})
	}

}

func bindTerm(c *gin.Context) (TermInput, error) {
	var input TermInput
	err := c.ShouldBindJSON(&input)
	return input, err
}

func existsTerm(id uuid.UUID) bool {
	var termFromDb models.Term
	err := models.DB.Model(&models.Term{}).Where("id = ?", id).Take(&termFromDb).Error
	if err != nil {
		return false
	}
	if termFromDb.ID != uuid.Nil {
		return true
	}
	return false
}
