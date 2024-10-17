package user

import (
	"barcation_be/handlers"
	"barcation_be/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserByIdController(c *gin.Context) {

	var user_id uint
	var err error

	if c.Query("id") != "" {
		uid, err := strconv.ParseUint(c.Query("id"), 10, 32)
		user_id = uint(uid)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
			return
		}
	} else {
		user_id, err = handlers.ExtractTokenById(c)
	}

	var u models.User

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	data, err := u.GetUserById(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	res := map[string]interface{}{
		"username": data.Username,
		"level":    data.Level,
		"email":    data.Email,
		"address":  data.Address,
		"phone":    data.Phone,
		"position": data.Position,
		"status":   data.Status,
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "id": data.ID, "user": res})
}
