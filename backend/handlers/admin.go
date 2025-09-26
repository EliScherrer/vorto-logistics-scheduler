package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateLoad handles load creation
func CreateLoad(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create load endpoint",
	})
}

// GetLoads handles retrieving all loads
func GetLoads(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get loads endpoint",
	})
}