package system

import (
	v1 "gin-api-learn/api/v1"
	"gin-api-learn/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(engine *gin.Engine) {
	userRouter := engine.Group("v1/user")
	{
		userRouter.POST("login", v1.Login)
		userRouter.POST("register", v1.Register)
	}
	tokenGroup := engine.Group("v1/user").Use(middleware.JWTAuthMiddleware())
	{
		tokenGroup.POST("/detail", v1.GetUser)
	}

}
