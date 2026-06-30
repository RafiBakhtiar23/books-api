package main

import (
	"log"

	"books-api/config"
	"books-api/models"
	"books-api/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	config.ConnectDB()

	// Auto migrate database
	config.DB.AutoMigrate(&models.Book{})
	router := gin.Default()
	routers.SetupRoutes(router)
	router.Run(":8080")
	log.Println("Books API is ready!")

}

