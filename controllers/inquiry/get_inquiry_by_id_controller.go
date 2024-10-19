package inquiry

import (
	"barcation_be/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetInquiryByIdController(c *gin.Context) {
	var inquiryId uint
	var i models.Inquiry

	if c.Query("id") != "" {
		uid, err := strconv.ParseUint(c.Query("id"), 10, 32)
		inquiryId = uint(uid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
			return
		}
	}

	data, err := i.GetInquiryById(inquiryId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	res := map[string]interface{}{
		"id":             i.ID,
		"user_name":      i.User.Username,
		"product_name":   i.Product.Name,
		"total_price":    i.TotalPrice,
		"total_quantity": i.TotalQuantity,
		"status":         i.Status,
		"created_at":     i.CreatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "id": data.ID, "product": res})
}
