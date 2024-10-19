package controllers

import (
	"barcation_be/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutController(c *gin.Context) {
	tokenString, _ := handlers.ExtractToken(c)

	// Add token to blacklist
	handlers.AddTokenToBlacklist(tokenString)

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "logout success"})
}
