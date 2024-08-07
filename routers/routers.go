package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phuongaz/forbo/routers/v1/feed"
	"github.com/phuongaz/forbo/routers/v1/user"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	setupRouter(router)
	return router
}

func setupRouter(router *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	router.Use(cors.New(corsConfig))

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			user.RegisterUserRouters(v1)
			feed.RegisterFeedRouters(v1)
		}
	}
}
