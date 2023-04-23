package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book id"})
		return
	}

	comment.BookID = uint(bookID)

	if err := initializers.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": comment})
}

func CreateBook(c *gin.Context) {
	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body"})
		return
	}

	cookie, err := c.Cookie("Authorization")
	book.UserID, err = getUserIDFromToken(cookie)
	book.Active = true

	result := initializers.DB.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func getUserIDFromToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Token error")
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	userID := token.Claims.(jwt.MapClaims)["sub"].(float64)

	return int(userID), nil
}
