package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	Description string
	Price       int
	Author      string
	Page        int
	Year        int
	Active      bool
	UserID      int
}
