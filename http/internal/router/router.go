package router

import (
	"github.com/gin-gonic/gin"
	user "github.com/sunflower10086/TikTok/http/internal/user/http"
	video "github.com/sunflower10086/TikTok/http/internal/video/http"
)

func Init(r *gin.Engine) {
	// noAuth
	noAuthRouter := r.Group("/douyin")
	{
		noAuthRouter.POST("/user/register", user.Register)
		noAuthRouter.POST("/user/login", user.Login)
		noAuthRouter.POST("/publish/action/", video.PublishAction)
		noAuthRouter.GET("/feed", video.GetFeedVideo)
	}

	// 加的auth.JwtAuthMiddleware()是一个登录验证的中间件
	// withAuthRouter := g.Group("/douyin", auth.JwtAuthMiddleware())
	// {
	// 	withAuthRouter.GET("/user/", u.GetInfoHandler) // 获取用户信息
	//  withAuthRouter.GET("/publish/list", video.GetPublishList)
	// }
}