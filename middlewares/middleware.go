package middlewares

import (
	"barcation_be/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := handlers.ValidateToken(c)

		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		if handlers.IsTokenBlacklisted(c) {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}
