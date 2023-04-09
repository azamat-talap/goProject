package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	BookID uint
	Text   string
}
