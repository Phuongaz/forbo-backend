package feed

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phuongaz/forbo/helper"
	"github.com/phuongaz/forbo/models"
)

func checkUserVaildMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newFeed models.FeedSkeleton
		if err := c.ShouldBindJSON(&newFeed); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if newFeed.UserID == "" {
			c.JSON(400, gin.H{"error": "UserID is required"})
			c.Abort()
			return
		}

		if newFeed.Content == "" {
			c.JSON(400, gin.H{"error": "Content is required"})
			c.Abort()
			return
		}

		userID := newFeed.UserID
		user, err := models.FindUserByID(userID)

		if err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		if user.UserID != newFeed.UserID {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func compareUserIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		feed, err := models.FindFeedByID(id)

		if err != nil {
			c.JSON(404, gin.H{"error": "Feed not found"})
			c.Abort()
			return
		}

		token := c.GetHeader("Authorization")
		claims, err := helper.GetClaimsFromToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			c.Abort()
			return
		}

		if claims.UserID != feed.UserID {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
