package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRouters(router *gin.RouterGroup) {
	user := router.Group("/user")
	{
		user.GET("/:id", getUser)
		user.POST("/register", registerUser)
		user.POST("/login", loginUser)
		user.POST("/follow", followUser)
		user.POST("/unfollow", unfollowUser)
		user.GET("/followers/:id", getFollowers)
		user.POST("/avatar", uploadAvatar)
		user.GET("/avatar/:id", getAvatar)
	}
}
