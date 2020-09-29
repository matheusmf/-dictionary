package routes

import (
	"dictionary/api/auth"
	"dictionary/api/controllers"

	"github.com/gin-gonic/gin"
)

// Initialize - define routes
func Initialize(router *gin.Engine) {
	// router.GET("/someGet", getting)
	// router.POST("/somePost", posting)
	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	// router.PATCH("/somePatch", patching)

	router.Use(auth.JwtFilter())

	router.GET("/terms", controllers.FindTerms)
	router.POST("/terms", controllers.CreateTerm)
}
