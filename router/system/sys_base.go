package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zhengpanone/gin-api-learn/api/v1"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("/api/v1")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.GET("captcha", baseApi.Captcha)
		baseRouter.GET("config", baseApi.GetConfig)
	}
	return baseRouter
}
