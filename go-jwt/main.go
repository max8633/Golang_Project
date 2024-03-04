package main

import (
	"github.com/gin-gonic/gin"
	"github.com/max8633/go-jwt/controllers"
	"github.com/max8633/go-jwt/initializers"
	"github.com/max8633/go-jwt/middlewares"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hi there")
	})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)

	r.Run()
}
