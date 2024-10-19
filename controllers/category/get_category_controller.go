package category

import (
	"barcation_be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategoryController(c *gin.Context) {
	var cat models.Category

	data, err := cat.GetCategory()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	res := make([]map[string]interface{}, 0)
	for _, v := range data {
		res = append(res, map[string]interface{}{
			"id":   v.ID,
			"name": v.Name,
			"icon": v.Icon,
		})
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "data": res})
}
