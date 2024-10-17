package user

import (
	"barcation_be/handlers"
	"barcation_be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type updatePasswordUserRequest struct {
	Password string `json:"password"`
}

func UpdatePasswordUserController(c *gin.Context) {

	var request updatePasswordUserRequest
	var u models.User
	user_id, err := handlers.ExtractTokenById(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	dataUser, err := u.GetUserById(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	err = dataUser.UpdatePassword(request.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "update password success"})

}
