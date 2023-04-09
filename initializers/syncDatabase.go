package initializers

import "github.com/tlpazmt/goProject/api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.Rating{})
	DB.AutoMigrate(&models.Comment{})
}
