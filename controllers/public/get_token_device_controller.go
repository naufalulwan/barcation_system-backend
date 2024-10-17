package public

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var tempVariable int = 0
var mu sync.Mutex

func GetTokenDeviceController(c *gin.Context) {
	newUUID := uuid.New()

	randVariable := rand.Intn(1000) + 1

	mu.Lock()
	tempVariable++
	mu.Unlock()

	deviceToken := fmt.Sprintf("%s-%v-%v", newUUID, tempVariable, randVariable)

	c.JSON(http.StatusOK, gin.H{"error": false, "code": http.StatusOK, "message": "success", "token": deviceToken})
}
