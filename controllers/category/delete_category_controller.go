package category

import (
	"barcation_be/handlers"
	"barcation_be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type deleteCategoryRequest struct {
	ID uint `json:"id"`
}

func DeleteCategoryController(c *gin.Context) {
	var cat models.Category
	var u models.User
	var request deleteCategoryRequest
	admin_id, err := handlers.ExtractTokenById(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	dataAdmin, err := u.GetUserById(admin_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if dataAdmin.Level != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": "only admin can delete category"})
		return
	}

	err = cat.DeleteCategory(request.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "delete category success"})

}
