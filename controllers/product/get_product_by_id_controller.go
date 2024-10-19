package product

import (
	"barcation_be/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProductByIdController(c *gin.Context) {
	var productId uint
	var p models.Product

	if c.Query("id") != "" {
		uid, err := strconv.ParseUint(c.Query("id"), 10, 32)
		productId = uint(uid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
			return
		}
	}

	data, err := p.GetProductById(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	res := map[string]interface{}{
		"name":          data.Name,
		"price":         data.Price,
		"stock":         data.Quantity,
		"status":        data.Status,
		"image":         data.Image,
		"desc":          data.Description,
		"category_name": data.Category.Name,
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "id": data.ID, "product": res})
}
