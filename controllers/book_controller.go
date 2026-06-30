package controllers

import (
	"books-api/config"
	"books-api/models"

	"github.com/gin-gonic/gin"
)

// Create/add
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	config.DB.Create(&book)
	c.JSON(200, book)
}

func GetAllBook(c *gin.Context) {
	var book []models.Book
	config.DB.Find(&book)
	c.JSON(200, book)
}

func GetByIdBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := config.DB.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"message": "Id Tidak Ditemukan", "id dicari": id})
		return
	}
	c.JSON(200, book)
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := config.DB.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"message": "Id Tidak Ditemukan"})
		return
	}
	c.ShouldBindJSON(&book)
	config.DB.Save(&book)
	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	var book models.Book

	if err := config.DB.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"message": "Id Tidak Ditemukan"})
		return
	}

	config.DB.Delete(&book)
	c.JSON(200, gin.H{"message": "book berhasil di hapus"})
}
