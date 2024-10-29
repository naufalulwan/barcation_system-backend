package product

import (
	"barcation_be/handlers"
	"barcation_be/helper"
	"barcation_be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createProductRequest struct {
	Name        string `form:"name"`
	Price       int    `form:"price"`
	Quantity    int    `form:"quantity"`
	Description string `form:"description"`
	CategoryID  int    `form:"category_id"`
}

func CreateProductController(c *gin.Context) {
	var u models.User
	var request createProductRequest

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
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "only admin can create product"})
		return
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image is required"})
		return
	}

	mimeType := file.Header.Get("Content-Type")
	if mimeType != "image/jpeg" && mimeType != "image/png" {
		helper.Logger.Debugf("Invalid image type: %s", mimeType)
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "Invalid image format. Only JPEG, PNG, and GIF are allowed."})
		return
	}

	p := models.Product{
		Name:        request.Name,
		Price:       request.Price,
		Quantity:    request.Quantity,
		Description: request.Description,
		Status:      true,
		CategoryID:  uint(request.CategoryID),
	}

	imagePath, err := helper.UploadImageHelper(file, "temp/product_image")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	p.Image = imagePath

	if err := p.SaveProduct(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := map[string]interface{}{
		"name":          p.Name,
		"price":         p.Price,
		"stock":         p.Quantity,
		"status":        p.Status,
		"image":         p.Image,
		"desc":          p.Description,
		"category_name": p.Category.Name,
	}

	c.JSON(http.StatusCreated, gin.H{"error": false, "code": http.StatusCreated, "message": "product created successfully", "id": p.ID, "product": res, "create_at": p.CreatedAt})
}
