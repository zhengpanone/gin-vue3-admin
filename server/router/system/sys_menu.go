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
	menuApi := v1.ApiGroupApp.SystemApiGroup.MenuApi
	{
		menuRouter.POST("addMenu", menuApi.AddMenu)              //新增菜单
		menuRouter.DELETE("deleteMenu", menuApi.DeleteMenuByIds) //删除菜单
		menuRouter.PUT("updateMenu", menuApi.UpdateMenu)         //更新菜单
		menuRouter.GET("getMenuById", menuApi.GetMenuById)       //获取菜单详情
	}

	{
		menuRouterWithoutRecord.POST("getMenuList", menuApi.GetMenuList) // 获取菜单树
	}
	return menuRouter
}
