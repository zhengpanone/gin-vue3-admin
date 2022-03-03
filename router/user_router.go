package router

import (
	v1 "gin-api-learn/api/v1"
	"gin-api-learn/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(engine *gin.Engine) {
	noLoginGroup := engine.Group("v1/user")
	{
		noLoginGroup.POST("login", v1.Login)
		noLoginGroup.POST("register", v1.Register)
	}
	tokenGroup := engine.Group("v1/user").Use(middleware.JWTAuthMiddleware())
	{
		tokenGroup.POST("/detail", v1.GetUser)
	}

}
