package controllers

import (
	"net/http"
	"strconv"

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

	db.Select("books.*, AVG(ratings.value) as rating").
		Joins("LEFT JOIN ratings ON ratings.book_id = books.id").
		Group("books.id").
		Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func SetBookRating(c *gin.Context) {
	bookID := c.Param("id")
	ratingValue, err := strconv.Atoi(c.PostForm("value"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rating value"})
		return
	}

	var book models.Book
	if err := initializers.DB.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
		return
	}

	rating := models.Rating{BookID: book.ID, Value: ratingValue}
	if err := initializers.DB.Create(&rating).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create rating"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": rating})
}
