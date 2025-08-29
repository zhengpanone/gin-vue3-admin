package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zhengpanone/gin-vue3-admin/api"
)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	roleRouter := Router.Group("v1/api/role")
	roleApi := v1.ApiGroupApp.RoleApi
	{
		roleRouter.POST("addRole", roleApi.CreateRole)      // 创建角色
		roleRouter.POST("pageRole", roleApi.PageRole)       // 分页获取角色
		roleRouter.DELETE("delete/:id", roleApi.DeleteRole) // 删除角色
		roleRouter.PUT("/update/:id", roleApi.UpdateRole)   // 更新角色
	}
	return roleRouter
}
