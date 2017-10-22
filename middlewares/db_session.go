// Package middlewares contains gin middlewares
// Usage: router.Use(middlewares.Connect)
package middlewares

import (
	"../db"
	"github.com/gin-gonic/gin"
)

// Connect middleware clones the database session for each request and
// makes the `db` object available for each handler
func Connect(conn *db.Connection) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Set example variable
		s := conn.Clone()
		defer conn.Close()

		c.Set("db", s)
		c.Next()
	}
}
