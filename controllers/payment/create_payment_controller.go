package payment

import (
	"barcation_be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreatePaymentRequest struct {
	UserID           int    `json:"user_id"`
	InquiryID        int    `json:"inquiry_id"`
	PaymentType      string `json:"payment_type"`
	PaymentReference string `json:"payment_reference"`
	PaymentSignature string `json:"payment_signature"`
	PaymentDate      string `json:"payment_date"`
}

func CreatePaymentController(c *gin.Context) {
	var request CreatePaymentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	payment := models.Payment{
		UserID:           uint(request.UserID),
		InquiryID:        uint(request.InquiryID),
		PaymentType:      request.PaymentType,
		PaymentStatus:    true,
		PaymentSignature: request.PaymentSignature,
		PaymentReference: request.PaymentReference,
		PaymentDate:      request.PaymentDate,
	}

	if err := payment.CreatePayment(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	res := map[string]interface{}{
		"id":                payment.ID,
		"user_name":         payment.User.Username,
		"product_name":      payment.Inquiry.Product.Name,
		"payment_reference": payment.PaymentReference,
		"payment_type":      payment.PaymentType,
		"payment_date":      payment.PaymentDate,
		"status":            payment.PaymentStatus,
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "data": res})
}
