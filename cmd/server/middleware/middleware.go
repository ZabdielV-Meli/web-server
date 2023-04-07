package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func VerificarToken() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenEnv := os.Getenv("TOKEN")
		tokenHeader := c.GetHeader("TOKEN")

		if tokenHeader != tokenEnv {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "El token no es el mismo"})
			return

		}

		c.Next()
	}
}
