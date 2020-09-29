package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

var DB *gorm.DB

// ConnectDatabase - initialize database connection
func ConnectDatabase() {
	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	DB, err = gorm.Open(os.Getenv("DB_DRIVER"), DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", os.Getenv("DB_DRIVER"))
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", os.Getenv("DB_DRIVER"))
	}

	DB.AutoMigrate(&Term{}, &RelatedTerm{})
}
