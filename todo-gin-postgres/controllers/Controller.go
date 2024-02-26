package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/max8633/todo-gin-postgres/initializers"
	"github.com/max8633/todo-gin-postgres/models"
)

func PostCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Title       string
		Description string
	}

	c.Bind(&body)

	// Create a post
	post := models.Todo{Title: body.Title, Description: body.Description}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Response with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Todo
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the posts
	var post models.Todo
	initializers.DB.First(&post, id)

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Title       string
		Description string
	}

	c.Bind(&body)

	// Find the post we're updating
	var post models.Todo
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Todo{
		Title:       body.Title,
		Description: body.Description,
	})

	// Response with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the posts
	initializers.DB.Delete(&models.Todo{}, id)

	// Response with it
	c.Status(200)
}
