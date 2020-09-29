package controllers

import (
	"dictionary/api/models"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

// FindTerms - Find all terms
// GET /terms
func FindTerms(c *gin.Context) {
	var terms []models.Term
	models.DB.Find(&terms)

	c.JSON(http.StatusOK, gin.H{"data": terms})
}

type CreateTermInput struct {
	Name         string `json:"name" binding:"required"`
	Pronuciation string `json:"pronuciation"`
	Definition   string `json:"definition"`
	Synonyms     string `json:"synonyms"`
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

	userID, _ := uuid.FromString("9c5151fd-8711-443c-a4f3-b977e5eae1e6")
	// Create term
	term := models.Term{Name: input.Name, Pronuciation: input.Pronuciation, Definition: input.Pronuciation, Synonyms: input.Synonyms, CreatedByID: userID}
	models.DB.Create(&term)

	c.JSON(http.StatusOK, gin.H{"data": term})
}

func bindTerm(c *gin.Context) (CreateTermInput, error) {
	var input CreateTermInput
	err := c.ShouldBindJSON(&input)
	return input, err
}
