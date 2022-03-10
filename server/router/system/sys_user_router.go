package system

import (
	v1 "gin-api-learn/api/v1"
	"gin-api-learn/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(engine *gin.Engine) {
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	userRouter := engine.Group("v1/user")
	{
		userRouter.POST("login", baseApi.Login)       // 用户登录
		userRouter.POST("register", baseApi.Register) // 用户注册
	}
	tokenGroup := engine.Group("v1/user").Use(middleware.JWTAuthMiddleware())
	{
		tokenGroup.POST("/detail", baseApi.GetUser)
	}

}
