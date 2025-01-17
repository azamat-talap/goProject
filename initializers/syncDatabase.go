package initializers

import "github.com/tlpazmt/goProject/api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.Rating{})
	DB.AutoMigrate(&models.Comment{})
	DB.AutoMigrate(&models.Role{})
	DB.AutoMigrate(&models.Purchase{})
	models.SeedRoles(DB)
}
