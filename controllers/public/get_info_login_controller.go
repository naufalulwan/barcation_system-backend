package public

import (
	"barcation_be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InfoLoginRequest struct {
	Username string `json:"username" binding:"required"`
	DeviceId string `json:"device_id" binding:"required"`
}

func GetInfoLoginController(c *gin.Context) {
	var request InfoLoginRequest
	var u models.User

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	data, err := u.GetUserByDeviceId(request.DeviceId, request.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if data.DeviceId != request.DeviceId {
		data.SaveLogin = false

		err = data.UpdateSaveInfoLogin()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "id device not match"})
		return
	}

	res := map[string]interface{}{
		"username":     data.Username,
		"device_id":    data.DeviceId,
		"is_save_info": data.SaveLogin,
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "id": data.ID, "user": res})
}
