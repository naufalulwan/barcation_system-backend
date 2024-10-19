package product

import (
	"barcation_be/handlers"
	"barcation_be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type updateProductRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Status      bool   `json:"status"`
	CategoryID  uint   `json:"category_id"`
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
		Status:      request.Status,
		CategoryID:  request.CategoryID,
	}

	err = p.UpdateProduct(request.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "update product success"})

}
