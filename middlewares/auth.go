package middlewares

import (
	"github.com/gin-gonic/gin"
)

// AuthMiddleware to authentificate and authorise user.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add your logic here

		c.Next()
	}
}
