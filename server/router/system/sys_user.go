package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zhengpanone/gin-vue3-admin/api/v1"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	userRouter := Router.Group("v1/api/admin")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("changePassword", baseApi.ChangePassword) //用户修改密码
		userRouter.POST("detail", baseApi.GetUser)
	}
	return userRouter

}
