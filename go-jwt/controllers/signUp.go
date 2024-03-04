package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/max8633/go-jwt/initializers"
	"github.com/max8633/go-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get Email/Username/Password off req body
	var body struct {
		Email    string
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})

		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})

		return
	}

	// Create user
	user := models.User{Email: body.Email, Username: body.Username, Password: string(hash)}

	if result := initializers.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})

		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{})
}
