package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zhengpanone/gin-vue3-admin/api"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("/v1/api")
	baseApi := v1.ApiGroupApp.BaseApi
	{
		baseRouter.POST("admin/login", baseApi.Login)       // 用户登录
		baseRouter.POST("admin/register", baseApi.Register) // 用户注册
		// baseRouter.POST("admin/logout", baseApi.Logout)     // 用户退出
		baseRouter.GET("captcha", baseApi.Captcha)  // 获取图片验证码
		baseRouter.GET("config", baseApi.GetConfig) // 获取配置

	}
	return baseRouter
}
