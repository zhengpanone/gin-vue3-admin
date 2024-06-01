package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zhengpanone/gin-vue3-admin/api/v1"
	"github.com/zhengpanone/gin-vue3-admin/middleware"
)

type MenuRouter struct{}

func (m *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("menu")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		menuRouter.POST("addBaseMenu", authorityMenuApi.AddBaseMenu) //新增菜单
	}

	{
		menuRouterWithoutRecord.POST("getMenu", authorityMenuApi.GetMenu)         // 获取菜单树
		menuRouterWithoutRecord.POST("getMenuList", authorityMenuApi.GetMenuList) // 获取菜单树
	}
	return menuRouter
}
