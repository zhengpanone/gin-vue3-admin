package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zhengpanone/gin-api-learn/api/v1"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.GET("config", baseApi.GetConfig)
	}
	return baseRouter
}
