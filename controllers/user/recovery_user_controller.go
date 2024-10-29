package user

import (
	"barcation_be/handlers"
	"barcation_be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecoveryUserRequest struct {
	Id     int  `json:"id" binding:"required"`
	Status bool `json:"status"`
}

func RecoveryUserController(c *gin.Context) {
	var req RecoveryUserRequest
	var u *models.User
	var userId *models.User

	adminId, err := handlers.ExtractTokenById(c)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "Invalid JSON input"})
		return
	}

	if adminId == 0 || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "Invalid admin credentials"})
		return
	}

	dataAdmin, err := u.GetUserById(adminId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if dataAdmin.Level != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "only admin can recovery user"})
		return
	}

	if req.Status == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "Status must be true for recovery"})
		return
	}

	if userId, err = u.GetUserById(uint(req.Id)); err != nil {
		userId, err = u.GetUserByDelete(uint(req.Id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": true, "code": http.StatusNotFound, "message": "User not found"})
			return
		}
	}

	if userId.Status == req.Status {
		c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "User is already active"})
		return
	}

	err = userId.RecoveryUser(uint(req.Id), req.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "user recovery success"})

}
