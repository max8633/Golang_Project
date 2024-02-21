package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// UserController represents a user-related controller
type UserController struct{}

func (uc *UserController) getUserInfo(c *gin.Context) {
	userID := c.Param("id")
	c.JSON(200, gin.H{"id": userID, "name": "Max Kuo", "email": "MaxKuo@example.com"})
}

// Middleware use to print related log
func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Request - Method: %s | Status: %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)
	}
}

// Middleware use to deal with authentication, but not really do auth features here!
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

func main() {
	//create a new gin router
	router := gin.Default()

	//setup custom middleware loggerMiddleware
	router.Use(loggerMiddleware())

	//define a route for root URL
	router.GET("/", func(c *gin.Context) {
		c.String(200, "welcome to gin!")
	})

	//define a route for /bye
	router.GET("/bye", func(c *gin.Context) {
		c.String(200, "good bye, have a good day!")
	})

	// Route with URL parameters
	router.GET("/User/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "Hi "+id)
	})

	//Route with query parameters
	router.GET("/search", func(c *gin.Context) {
		query := c.DefaultQuery("q", "default-value")
		c.String(200, "Search query: "+query)
	})

	// Public routes (no authentication required)
	public := router.Group("/public")
	{
		public.GET("/info", func(c *gin.Context) {
			c.String(200, "Public Information")
		})
		public.GET("/products", func(c *gin.Context) {
			c.String(200, "Public Products List")
		})
	}

	// Private routes (require authentication)
	private := router.Group("/private")
	private.Use(authMiddleware())
	{
		private.GET("/data", func(c *gin.Context) {
			c.String(200, "Private data accessible after authentication")
		})
		private.POST("/create", func(c *gin.Context) {
			c.String(200, "Create a new resource")
		})
	}

	/*define UserController struct and handle related task and logic
	Separating business logic into controllers
	makes the codebase cleaner and more organized, improving readability and maintainability. */
	uc := &UserController{}
	router.GET("/users/:id", uc.getUserInfo)

	//run the server on port 8080
	router.Run(":8080")
}
