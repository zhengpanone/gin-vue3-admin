package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-api-learn/api/v1/system"
	"github.com/zhengpanone/gin-api-learn/middleware"
)

func InitUserRouter(engine *gin.Engine) {
	userRouter := engine.Group("v1/user")

	{

		userRouter.POST("register", system.Register)             // 用户注册
		userRouter.POST("changePassword", system.ChangePassword) //用户修改密码

	}
	tokenGroup := engine.Group("v1/user").Use(middleware.JWTAuthMiddleware())
	{
		tokenGroup.POST("/detail", system.GetUser)
	}

}
