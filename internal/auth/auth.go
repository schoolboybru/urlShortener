package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if strings.HasPrefix(authorization, "Bearer ") {
			splits := strings.Split(authorization, " ")
			if len(splits) != 2 {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Token is invalid"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Missing Token"})
		}
		c.Next()
	}
}
