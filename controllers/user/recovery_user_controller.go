package user

import (
	"barcation_be/handlers"
	"barcation_be/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecoveryUserRequest struct {
	Id     int  `json:"id" binding:"required"`
	Status bool `json:"status"`
}

func RecoveryUserController(c *gin.Context) {
	var req RecoveryUserRequest
	var u models.User
	admin_id, err := handlers.ExtractTokenById(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	dataAdmin, err := u.GetUserById(admin_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if dataAdmin.Level != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "only admin can recovery user"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if req.Status == false {
		fmt.Printf("%v", req.Status)
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "status must be true"})
		return
	}

	err = u.RecoveryUser(uint(req.Id), req.Status)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "user recovery success"})

}
