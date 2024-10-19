package inquiry

import (
	"barcation_be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createInquiryRequest struct {
	TotalQuantity int `json:"total_quantity"`
	TotalPrice    int `json:"total_price"`
	ProductID     int `json:"product_id"`
	UserID        int `json:"user_id"`
}

func CreateInquiryController(c *gin.Context) {
	var request createInquiryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	inquiry := models.Inquiry{
		TotalQuantity: request.TotalQuantity,
		TotalPrice:    request.TotalPrice,
		Status:        true,
		ProductID:     uint(request.ProductID),
		UserID:        uint(request.UserID),
	}

	if err := inquiry.CreateInquiry(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	res := map[string]interface{}{
		"total_quantity": inquiry.TotalQuantity,
		"total_price":    inquiry.TotalPrice,
		"product_name":   inquiry.Product.Name,
		"product_price":  inquiry.Product.Price,
		"user_name":      inquiry.User.Username,
		"status":         inquiry.Status,
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "data": res})
}
