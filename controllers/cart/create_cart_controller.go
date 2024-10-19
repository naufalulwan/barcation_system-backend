package cart

import (
	"barcation_be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createCartRequest struct {
	Quantity  int `json:"quantity"`
	ProductID int `json:"product_id"`
	UserID    int `json:"user_id"`
}

func CreateCartController(c *gin.Context) {
	var p models.Product
	var request createCartRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	productData, err := p.GetProductById(uint(request.ProductID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	cart := models.Cart{
		Quantity:  request.Quantity,
		Total:     productData.Price * request.Quantity,
		ProductID: uint(request.ProductID),
		UserID:    uint(request.UserID),
	}

	if err := cart.AddCart(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	res := map[string]interface{}{
		"quantity": cart.Quantity,
		"total":    cart.Total,
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "data": res})
}
