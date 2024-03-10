package initializers

import (
	"github.com/max8633/Golang_Project/todo-gin-postgres/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.Todo{})
}
