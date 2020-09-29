package api

import (
	"dictionary/api/models"
	"dictionary/api/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

// Run initializer routes and server
func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	models.ConnectDatabase()

	router := gin.Default()

	routes.Initialize(router)

	router.Run()

}
