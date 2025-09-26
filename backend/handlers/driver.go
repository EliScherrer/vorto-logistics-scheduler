package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StartShift handles driver shift start
func StartShift(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Start shift endpoint",
	})
}

// EndShift handles driver shift end
func EndShift(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "End shift endpoint",
	})
}

// GetDriverStatus handles getting driver status
func GetDriverStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get driver status endpoint",
	})
}