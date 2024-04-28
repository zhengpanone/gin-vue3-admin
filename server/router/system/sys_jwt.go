package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zhengpanone/gin-vue3-admin/api/v1"
)

type JwtRouter struct{}

func (j *JwtRouter) InitJwtRouter(Router *gin.RouterGroup) {
	jwtRouter := Router.Group("jwt")
	jwtApi := v1.ApiGroupApp.SystemApiGroup.JwtApi
	{
		jwtRouter.POST("jwtInBlacklist", jwtApi.JsonInBlacklist) // jwt加入黑名单
	}
}
