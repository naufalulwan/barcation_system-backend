package inquiry

import (
	"barcation_be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInquiryController(c *gin.Context) {
	var i models.Inquiry

	data, err := i.GetInquiry()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	res := make([]map[string]interface{}, 0)
	for _, v := range data {
		res = append(res, map[string]interface{}{
			"id":             v.ID,
			"user_name":      v.User.Username,
			"product_name":   v.Product.Name,
			"total_price":    v.TotalPrice,
			"total_quantity": v.TotalQuantity,
			"status":         v.Status,
			"created_at":     v.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "data": res})
}
