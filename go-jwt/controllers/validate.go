package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {

	c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": "login",
	})
}
