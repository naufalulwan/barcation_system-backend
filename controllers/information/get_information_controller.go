package information

import (
	"barcation_be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInformationController(c *gin.Context) {
	var typeInformation string
	var i models.Information

	if c.Query("type") != "" {
		typeInformation = c.Query("type")
	}

	data, err := i.GetInformation(typeInformation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	res := make([]map[string]interface{}, 0)
	for _, v := range data {
		typeInformation = v.Type
		res = append(res, map[string]interface{}{
			"id":      v.ID,
			"title":   v.Title,
			"message": v.Message,
			"image":   v.Image,
		})
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "type": typeInformation, "product": res})
}
