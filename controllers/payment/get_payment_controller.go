package payment

import (
	"barcation_be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPaymentController(c *gin.Context) {
	var pay models.Payment

	data, err := pay.GetPayment()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	res := make([]map[string]interface{}, 0)
	for _, v := range data {
		res = append(res, map[string]interface{}{
			"id":                v.ID,
			"user_name":         v.User.Username,
			"product_name":      v.Inquiry.Product.Name,
			"payment_reference": v.PaymentReference,
			"payment_type":      v.PaymentType,
			"payment_date":      v.PaymentDate,
			"total_price":       v.Inquiry.TotalPrice,
			"total_quantity":    v.Inquiry.TotalQuantity,
			"status":            v.PaymentStatus,
		})
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "data": res})
}
