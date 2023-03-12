package initializers

import "github.com/tlpazmt/goProject/api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}
