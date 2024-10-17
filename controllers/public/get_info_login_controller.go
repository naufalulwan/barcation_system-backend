package public

import (
	"barcation_be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InfoLoginRequest struct {
	Username string `json:"username" binding:"required"`
	DeviceId string `json:"device_id" binding:"required"`
	Ssn      string `json:"ssn" binding:"required"`
}

func GetInfoLoginController(c *gin.Context) {
	var request InfoLoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	var u models.User

	data, err := u.GetUserByDeviceId(request.DeviceId, request.Username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if data.Ssn != request.Ssn {
		data.SaveLogin = false

		err = data.UpdateSaveInfoLogin()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "ssn not match"})
		return
	}

	res := map[string]interface{}{
		"username":     data.Username,
		"device_id":    data.DeviceId,
		"ssn":          data.Ssn,
		"is_save_info": data.SaveLogin,
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "id": data.ID, "user": res})
}
