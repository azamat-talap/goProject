package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tlpazmt/goProject/api/models"
	"github.com/tlpazmt/goProject/initializers"
)

func GetBooks(c *gin.Context) {
	title := c.Query("title")
	var books []models.Book
	initializers.DB.Where("title LIKE ?", "%"+title+"%").Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}
