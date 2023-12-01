package controllers

import (
	"github.com/f1rezy/book-store/internal/models"
	"github.com/f1rezy/book-store/pkg/db/postgresql"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	postgresql.DB.Create(&book)

	c.JSON(200, gin.H{"success": "book created"})
}

func GetAllBooks(c *gin.Context) {
	var books []models.Book

	postgresql.DB.Find(&books)

	c.JSON(200, books)
}

func GetBook(c *gin.Context) {
	pk := c.Param("pk")

	var book models.Book

	postgresql.DB.Where("id = ?", pk).First(&book)

	if book.ID == 0 {
		c.JSON(404, gin.H{"error": "book not found"})
		return
	}

	c.JSON(200, book)
}
