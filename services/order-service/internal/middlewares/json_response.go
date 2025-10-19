package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			status, _ := c.Get("status")
			message, _ := c.Get("message")

			statusCode, ok := status.(int)
			if !ok || statusCode == 0 {
				statusCode = http.StatusInternalServerError
			}

			msg, ok := message.(string)
			if !ok || msg == "" {
				msg = "Internal server error"
			}

			c.JSON(statusCode, gin.H{
				"success": false,
				"message": msg,
				"error":   err.Error(),
			})
			return
		}
	}
}
