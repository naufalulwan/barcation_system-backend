package payment

import (
	"barcation_be/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetPaymentByIdController(c *gin.Context) {
	var paymentId uint
	var pay models.Payment

	if c.Query("id") != "" {
		uid, err := strconv.ParseUint(c.Query("id"), 10, 32)
		paymentId = uint(uid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
			return
		}
	}

	data, err := pay.GetPaymentById(paymentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	res := map[string]interface{}{
		"id":                pay.ID,
		"user_name":         pay.User.Username,
		"product_name":      pay.Inquiry.Product.Name,
		"payment_reference": pay.PaymentReference,
		"payment_type":      pay.PaymentType,
		"payment_date":      pay.PaymentDate,
		"total_price":       pay.Inquiry.TotalPrice,
		"total_quantity":    pay.Inquiry.TotalQuantity,
		"status":            pay.PaymentStatus,
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "id": data.ID, "product": res})
}
