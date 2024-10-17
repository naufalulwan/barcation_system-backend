package product

import (
	"barcation_be/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProductByIdController(c *gin.Context) {
	var product_id uint

	if c.Query("id") != "" {
		uid, err := strconv.ParseUint(c.Query("id"), 10, 32)
		product_id = uint(uid)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
			return
		}
	}

	var p models.Product

	data, err := p.GetProductById(product_id)

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
