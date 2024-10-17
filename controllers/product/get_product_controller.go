package product

import (
	"barcation_be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductController(c *gin.Context) {

	var p models.Product

	data, err := p.GetProduct()

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	res := make([]map[string]interface{}, 0)
	for _, v := range data {
		res = append(res, map[string]interface{}{
			"id":            v.ID,
			"name":          v.Name,
			"price":         v.Price,
			"stock":         v.Quantity,
			"status":        v.Status,
			"image":         v.Image,
			"desc":          v.Description,
			"category_name": v.Category.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "data": res})

}
