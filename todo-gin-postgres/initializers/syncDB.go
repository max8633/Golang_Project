package initializers

import (
	"github.com/max8633/todo-gin-postgres/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.Todo{})
}
