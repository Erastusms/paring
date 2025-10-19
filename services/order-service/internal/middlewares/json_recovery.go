package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONRecovery menangkap panic dan kembalikan JSON error
func JSONRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "Internal server error",
					"error":   r,
				})
			}
		}()
		c.Next()
	}
}
