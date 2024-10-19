package user

import (
	"barcation_be/handlers"
	"barcation_be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type updateUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Position string `json:"position"`
}

func UpdateUserController(c *gin.Context) {

	user_id, err := handlers.ExtractTokenById(c)

	var u models.User

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	data, err := u.GetUserById(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	var request updateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	u.Email = request.Email
	u.Address = request.Address
	u.Phone = request.Phone
	u.Position = request.Position

	err = u.UpdateUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	res := map[string]interface{}{
		"username": data.Username,
		"level":    data.Level,
		"email":    u.Email,
	}

	optionalRes := map[string]interface{}{
		"address":  u.Address,
		"phone":    u.Phone,
		"position": u.Position,
	}

	for k, v := range optionalRes {
		if v != "" {
			res[k] = v
		}
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "update success", "id": data.ID, "user": res, "updated_at": u.UpdatedAt})

}
