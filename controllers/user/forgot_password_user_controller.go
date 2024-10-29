package user

import (
	"barcation_be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ForgotPasswordUserRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}

func ForgotPasswordUserController(c *gin.Context) {
	var request ForgotPasswordUserRequest
	var u models.User

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	dataUser, err := u.GetUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "Email is valid"})
	}

	if request.NewPassword != "" {
		err = dataUser.UpdatePassword(request.NewPassword)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "forgot password success"})

	}

}
