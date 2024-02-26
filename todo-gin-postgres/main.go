package main

import (
	"github.com/gin-gonic/gin"
	"github.com/max8633/todo-gin-postgres/controllers"
	"github.com/max8633/todo-gin-postgres/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnetToDB()
	initializers.SyncDatabase()
}

func main() {

	router := gin.Default()

	router.POST("/todo", controllers.PostCreate)
	router.PUT("/todo/:id", controllers.PostsUpdate)
	router.GET("/todo", controllers.PostsIndex)
	router.GET("/todo/:id", controllers.PostsShow)
	router.DELETE("/todo/:id", controllers.PostsDelete)

	router.Run()
}
