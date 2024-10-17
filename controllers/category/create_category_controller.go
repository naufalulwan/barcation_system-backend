package category

import (
	"barcation_be/handlers"
	"barcation_be/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

type createCategoryRequest struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func CreateCategoryController(c *gin.Context) {
	admin_id, err := handlers.ExtractTokenById(c)
	var cat models.Category
	var u models.User
	var request createCategoryRequest

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
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "only admin can create category"})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	cat.Name = request.Name
	cat.Icon = request.Icon

	err = cat.SaveCategory()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "create category success"})

}
