package information

import (
	"barcation_be/handlers"
	"barcation_be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createInformationRequest struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Message string `json:"message"`
	Image   string `json:"image"`
}

func CreateInformationController(c *gin.Context) {
	var i models.Information
	var u models.User
	var request createInformationRequest

	adminId, err := handlers.ExtractTokenById(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	dataAdmin, err := u.GetUserById(adminId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if dataAdmin.Level != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "only admin can create information"})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	i.Type = request.Type
	i.Title = request.Title
	i.Message = request.Message
	i.Image = request.Image

	err = i.CreateInformation()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "create information success"})
}
