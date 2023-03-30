package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tlpazmt/goProject/api/models"
	"github.com/tlpazmt/goProject/initializers"
)

func GetBooks(c *gin.Context) {
	title := c.Query("title")
	minPrice := c.Query("min_price")
	maxPrice := c.Query("max_price")
	minRating := c.Query("min_rating")

	var books []models.Book
	db := initializers.DB

	if title != "" {
		db = db.Where("title LIKE ?", "%"+title+"%")
	}
	if minPrice != "" {
		db = db.Where("price >= ?", minPrice)
	}
	if maxPrice != "" {
		db = db.Where("price <= ?", maxPrice)
	}
	if minRating != "" {
		db = db.Where("rating >= ?", minRating)
	}

	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}
