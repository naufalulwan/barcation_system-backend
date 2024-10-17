package user

import (
	"barcation_be/handlers"
	"barcation_be/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type updateLevelUserRequest struct {
	Id    uint   `json:"id" binding:"required"`
	Level string `json:"level" binding:"required"`
}

func UpdateLevelUserController(c *gin.Context) {
	token_id, err := handlers.ExtractTokenById(c)

	var userAdmin *models.User
	var userUser *models.User

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	dataAdmin, err := userAdmin.GetUserById(token_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if dataAdmin.Level != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "only admin can update level user"})
		return
	}

	var request updateLevelUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	dataUser, err := userUser.GetUserById(request.Id)

	fmt.Println(userUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	dataUser.Level = request.Level

	err = dataUser.UpdateUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	res := map[string]interface{}{
		"username": dataUser.Username,
		"level":    dataUser.Level,
		"email":    dataUser.Email,
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "update success", "id": dataUser.ID, "user": res, "updated_at": dataUser.UpdatedAt})
}
