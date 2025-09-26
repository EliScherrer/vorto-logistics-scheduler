package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminLogin handles admin authentication
func AdminLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Admin login endpoint",
	})
}

// DriverLogin handles driver authentication
func DriverLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Driver login endpoint",
	})
}