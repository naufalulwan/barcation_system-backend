package user

import (
	"barcation_be/handlers"
	"barcation_be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserController(c *gin.Context) {
	user_id, err := handlers.ExtractTokenById(c)
	token, _ := handlers.ExtractToken(c)
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

	data.Status = false

	err = data.DeleteUser()
	handlers.AddTokenToBlacklist(token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "user deleted successfully, the account will be deleted in 7 days"})
}
