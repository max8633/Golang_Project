package initializers

import "github.com/max8633/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
