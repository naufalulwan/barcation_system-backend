package product

import (
	"barcation_be/models"
	"barcation_be/handlers"

	"net/http"

	"github.com/gin-gonic/gin"
)

type createProductRequest struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
	Image       string `json:"image"`
	CategoryID  uint   `json:"category_id"`
}

func CreateProductController(c *gin.Context) {

	var u models.User
	var request createProductRequest
	admin_id, err := handlers.ExtractTokenById(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	dataAdmin, err := u.GetUserById(admin_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if dataAdmin.Level != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "only admin can create product"})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	p := models.Product{
		Name:        request.Name,
		Price:       request.Price,
		Quantity:    request.Quantity,
		Description: request.Description,
		Image:       request.Image,
		Status:      true,
		CategoryID:  request.CategoryID,
	}

	err = p.SaveProduct()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
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
