package controllers

import (
	"dictionary/api/auth"
	"dictionary/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindTerms - Find all terms
// GET /terms
func FindTerms(c *gin.Context) {
	var terms []models.Term
	models.DB.Preload("RelatedTerms").Find(&terms)

	c.JSON(http.StatusOK, gin.H{"data": terms})
}

// CreateTermInput - DTO
type CreateTermInput struct {
	Name         string               `json:"name" binding:"required"`
	Pronuciation string               `json:"pronuciation"`
	Definition   string               `json:"definition"`
	Synonyms     string               `json:"synonyms"`
	RelatedTerms []models.RelatedTerm `json:"related_terms"`
}

// CreateTerm - Create new book
// POST /books
func CreateTerm(c *gin.Context) {
	// Validate input
	input, err := bindTerm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := auth.GetLoggedUserID(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

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

func bindTerm(c *gin.Context) (CreateTermInput, error) {
	var input CreateTermInput
	err := c.ShouldBindJSON(&input)
	return input, err
}
