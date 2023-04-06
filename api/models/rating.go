package models

import (
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	BookID uint
	Value  int
}
