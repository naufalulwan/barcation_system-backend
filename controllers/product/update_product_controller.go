package product

import (
	"barcation_be/handlers"
	"barcation_be/helper"
	"barcation_be/models"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type updateProductRequest struct {
	ID          uint   `form:"id"`
	Name        string `form:"name"`
	Price       int    `form:"price"`
	Quantity    int    `form:"quantity"`
	Description string `form:"description"`
	Status      bool   `form:"status"`
	CategoryID  int    `form:"category_id"`
}

func UpdateProductController(c *gin.Context) {
	var u models.User
	var request updateProductRequest

	adminId, err := handlers.ExtractTokenById(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	dataAdmin, err := u.GetUserById(adminId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if dataAdmin.Level != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "only admin can update product"})
		return
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	p := models.Product{
		Name:        request.Name,
		Price:       request.Price,
		Quantity:    request.Quantity,
		Description: request.Description,
		Status:      request.Status,
		CategoryID:  uint(request.CategoryID),
	}

	if file, err := c.FormFile("image"); err == nil {
		mimeType := file.Header.Get("Content-Type")

		if mimeType != "image/jpeg" && mimeType != "image/png" {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "Invalid image format. Only JPEG, PNG, and GIF are allowed."})
			return
		}

		data, err := p.GetProductById(request.ID)
		if err != nil {
			return
		}

		oldImagePath := data.Image
		reOldImagePath := strings.ReplaceAll(oldImagePath, "/", "\\")

		if _, err := os.Stat(reOldImagePath); err == nil {
			err := os.Remove(reOldImagePath)
			if err != nil {
				helper.Logger.Debugf("Error deleting old image: %v", err)
			} else {
				helper.Logger.Infof("Successfully deleted old image: %s", reOldImagePath)
			}
		} else if os.IsNotExist(err) {
			helper.Logger.Debugf("Old image does not exist: %s", reOldImagePath)
		} else {
			helper.Logger.Debugf("Error checking old image: %v", err)
		}

		imagePath, err := helper.UploadImageHelper(file, "temp/product_image")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		p.Image = imagePath
	}

	err = p.UpdateProduct(request.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "update product success"})
}
