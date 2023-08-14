package routes

import (
	"posthub/controller"
	"posthub/logger"
	"posthub/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(mode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	rg := r.Group("/api/v1")
	rg.POST("/signup", controller.SignUp)
	rg.POST("/login", controller.Login)
	rg.Use(middleware.JWTAuthMiddleware)

	{
		rg.GET("/community", controller.CommunityList)
		rg.GET("/community/:id", controller.Community)
		rg.POST("/post", controller.CreatePost)
		rg.GET("/post/:id", controller.PostDetail)
		rg.GET("/post", controller.Postlist)
		rg.POST("/vote", controller.Vote)
		rg.GET("/posts", controller.Posts)
	}
	return r
}
