package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func ServeImageHelper(c *gin.Context) {
	filename := c.Param("filename")
	filePath := filepath.Join("temp/product_image", filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}

	ext := filepath.Ext(filePath)
	switch ext {
	case ".jpg", ".jpeg":
		c.Header("Content-Type", "image/jpeg")
	case ".png":
		c.Header("Content-Type", "image/png")
	case ".gif":
		c.Header("Content-Type", "image/gif")
	default:
		c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "unsupported media type"})
		return
	}

	c.File(filePath)
}
