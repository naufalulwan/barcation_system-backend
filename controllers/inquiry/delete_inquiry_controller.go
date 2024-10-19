package inquiry

import (
	"barcation_be/handlers"
	"barcation_be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteInquiryRequest struct {
	ID uint `json:"id"`
}

func DeleteInquiryController(c *gin.Context) {
	var inquiry models.Inquiry
	var u models.User
	var request DeleteInquiryRequest

	userId, err := handlers.ExtractTokenById(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	dataUser, err := u.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if dataUser.Level != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "only admin can delete inquiry"})
		return
	}

	err = inquiry.DeleteInquiry(request.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "delete cart success"})
}
