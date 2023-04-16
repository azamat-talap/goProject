package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string `gorm:"not null;unique"`
}

func SeedRoles(db *gorm.DB) error {
	roles := []Role{
		{Name: "admin"},
		{Name: "client"},
		{Name: "seller"},
	}
	for _, role := range roles {
		if err := db.Create(&role).Error; err != nil {
			return err
		}
	}
	return nil
}
