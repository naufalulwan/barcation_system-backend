package controllers

import (
	"barcation_be/handlers"
	"barcation_be/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	DeviceId    string `json:"device_id" binding:"required"`
	DeviceToken string `json:"device_token" binding:"required"`
	IsSaveInfo  bool   `json:"is_save_info"`
}

func LoginController(c *gin.Context) {
	var req loginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "Terjadi kesalahan saat menerima data login, silahkan cek kembali."})
		return
	}

	u := models.User{
		Username:    req.Username,
		Password:    req.Password,
		DeviceId:    req.DeviceId,
		SaveLogin:   req.IsSaveInfo,
		DeviceToken: req.DeviceToken,
	}

	accessToken, refreshToken, user, err := handlers.AuthHandler(u.Username, u.Password, u.DeviceId, u.DeviceToken, u.SaveLogin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "code": http.StatusUnauthorized, "message": err.Error()})
		return
	}

	res := map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
		"level":    user.Level,
		"status":   user.Status,
	}

	tokenRes := map[string]interface{}{
		"token":         accessToken,
		"refresh_token": refreshToken,
		"type":          "Bearer",
	}

	c.Writer.Header().Set("Authorization", "Bearer "+accessToken)
	c.Writer.Header().Set("Refresh-Token", refreshToken)

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "login success", "id": user.ID, "user": res, "authorization": tokenRes})
}
