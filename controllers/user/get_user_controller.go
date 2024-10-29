package user

import (
	"barcation_be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserController(c *gin.Context) {
	var u models.User

	data, err := u.GetUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	res := make([]map[string]interface{}, 0)
	for _, v := range data {
		res = append(res, map[string]interface{}{
			"id":       v.ID,
			"username": v.Username,
			"email":    v.Email,
			"level":    v.Level,
		})
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "data": res})
}
