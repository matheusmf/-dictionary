package controllers

import (
	"dictionary/api/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func GetLoggedUserID(c *gin.Context) uuid.UUID {
	userID, err := auth.GetLoggedUserID(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	return userID
}
