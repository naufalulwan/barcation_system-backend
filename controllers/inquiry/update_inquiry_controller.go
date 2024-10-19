package inquiry

import (
	"barcation_be/handlers"
	"barcation_be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type updateInquiryRequest struct {
	ID            uint `json:"id"`
	TotalPrice    int  `json:"total_price"`
	TotalQuantity int  `json:"total_quantity"`
	Status        bool `json:"status"`
	UserID        uint `json:"user_id"`
	ProductID     uint `json:"product_id"`
}

func UpdateInquiryController(c *gin.Context) {
	var u models.User
	var request updateInquiryRequest

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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	i := models.Inquiry{
		TotalPrice:    request.TotalPrice,
		TotalQuantity: request.TotalQuantity,
		Status:        request.Status,
		UserID:        request.UserID,
		ProductID:     request.ProductID,
	}

	err = i.UpdateInquiry(request.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "update inquiry success"})
}
