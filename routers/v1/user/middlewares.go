package user

import (
	"github.com/gin-gonic/gin"
	"github.com/phuongaz/forbo/helper"
)

func compareUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims, err := helper.GetClaimsFromToken(token)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if claims.UserID != c.Param("id") {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
