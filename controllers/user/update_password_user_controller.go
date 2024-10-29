package user

import (
	"barcation_be/handlers"
	"barcation_be/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type updatePasswordUserRequest struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

func UpdatePasswordUserController(c *gin.Context) {
	var request updatePasswordUserRequest
	var u models.User

	userId, err := handlers.ExtractTokenById(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	dataUser, err := u.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(request.OldPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password lama tidak sesuai"})
		return
	}

	err = dataUser.UpdatePassword(request.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "update password success"})
}
