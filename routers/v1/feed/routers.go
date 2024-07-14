package feed

import (
	"github.com/gin-gonic/gin"
)

func RegisterFeedRouters(router *gin.RouterGroup) {
	user := router.Group("/feed")
	{
		user.GET("/:id", getFeedByID)                                //Get Feed
		user.GET("/user/:id", getFeedsByUserID)                      //Get Feed by User
		user.POST("/create", checkUserVaildMiddleware(), createFeed) //Create Feed
		user.PUT("/:id", compareUserIDMiddleware(), updateFeed)      //Post Feed
		user.DELETE("/:id", compareUserIDMiddleware(), deleteFeed)   //Delete Feed
	}
}
