package models

import (
	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	BookID  uint
	UserID  int
	Address string
}
