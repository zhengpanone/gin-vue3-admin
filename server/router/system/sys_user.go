package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-vue3-admin/api/v1/system"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	userRouter := Router.Group("v1/api/admin")

	{
		userRouter.POST("changePassword", system.ChangePassword) //用户修改密码
		userRouter.POST("detail", system.GetUser)
	}
	return userRouter

}
