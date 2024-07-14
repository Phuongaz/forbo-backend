package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phuongaz/forbo/helper"
)

func AdminAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
			c.Abort()
			return
		}
		claims, err := helper.GetClaimsFromToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			c.Abort()
			return
		}
		if !claims.IsAdmin() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Permission denied"})
			c.Abort()
			return
		}
		c.Next()
	}
}
