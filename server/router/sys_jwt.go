package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zhengpanone/gin-vue3-admin/api"
)

type JwtRouter struct{}

func (j *JwtRouter) InitJwtRouter(Router *gin.RouterGroup) {
	jwtRouter := Router.Group("jwt")
	jwtApi := v1.ApiGroupApp.JwtApi
	{
		jwtRouter.POST("jwtInBlacklist", jwtApi.JsonInBlacklist) // jwt加入黑名单
	}
}
