package payment

import (
	"barcation_be/handlers"
	"barcation_be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeletePaymentRequest struct {
	ID uint `json:"id"`
}

func DeletePaymentController(c *gin.Context) {
	var payment models.Payment
	var u models.User
	var request DeletePaymentRequest

	userId, err := handlers.ExtractTokenById(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	dataUser, err := u.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if dataUser.Level != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "only admin can delete payment"})
		return
	}

	err = payment.DeletePayment(request.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "delete payment success"})
}
