package cart

import (
	"barcation_be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCartController(c *gin.Context) {
	var cart models.Cart

	data, err := cart.GetCart()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	res := make([]map[string]interface{}, 0)
	for _, v := range data {
		res = append(res, map[string]interface{}{
			"id": v.ID,
			"product_data": map[string]interface{}{
				"id":    v.Product.ID,
				"name":  v.Product.Name,
				"image": v.Product.Image,
				"price": v.Product.Price,
				"stock": v.Product.Quantity,
			},
			"user_data": map[string]interface{}{
				"id":       v.User.ID,
				"username": v.User.Username,
				"email":    v.User.Email,
			},
			"quantity": v.Quantity,
			"total":    v.Total,
		})
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "data": res})
}
