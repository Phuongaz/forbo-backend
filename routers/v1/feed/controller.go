package feed

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/phuongaz/forbo/models"
)

func getFeedByID(c *gin.Context) {
	id := c.Param("id")
	feed, err := models.FindFeedByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Feed not found"})
		return
	}

	c.JSON(200, feed)
}

func getFeedsByUserID(c *gin.Context) {
	userID := c.Param("id")
	feeds, err := models.FindFeedsByUserID(userID)

	if err != nil {
		c.JSON(404, gin.H{"error": "Feed not found"})
		return
	}

	c.JSON(200, feeds)
}

func createFeed(c *gin.Context) {
	var newFeed models.FeedSkeleton
	if err := c.ShouldBindJSON(&newFeed); err != nil {
		log.Default().Println("Line 36,")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	feed := newFeed.ToFeed()

	if err := feed.Create(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Feed created successfully"})
}

func updateFeed(c *gin.Context) {
	id := c.Param("id")
	feed, err := models.FindFeedByID(id)

	if err != nil {
		c.JSON(404, gin.H{"error": "Feed not found"})
		return
	}

	if err := c.ShouldBindJSON(&feed); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := feed.Update(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Feed updated successfully"})
}

func deleteFeed(c *gin.Context) {
	id := c.Param("id")

	feed, err := models.FindFeedByID(id)

	if feed.UserID != c.GetString("userID") {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	if err != nil {
		c.JSON(404, gin.H{"error": "Feed not found"})
		return
	}

	err = feed.Delete()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Feed deleted successfully"})
}
