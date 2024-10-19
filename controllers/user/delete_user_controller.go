package user

import (
	"barcation_be/handlers"
	"barcation_be/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteUserController(c *gin.Context) {
	var userId uint
	var token string
	var u models.User
	var err error

	if c.Query("id") != "" {
		uid, err := strconv.ParseUint(c.Query("id"), 10, 32)
		userId = uint(uid)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
			return
		}
	} else {
		userId, err = handlers.ExtractTokenById(c)
		token, _ = handlers.ExtractToken(c)
	}

	data, err := u.GetUserById(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	data.Status = false

	err = data.DeleteUser()

	if token != "" {
		handlers.AddTokenToBlacklist(token)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "user deleted successfully, the account will be deleted in 7 days"})
}
