package api

import (
	"github.com/zhengpanone/gin-vue3-admin/service"
)

var (
	userService = service.ServiceGroupApp.UserService
	jwtService  = service.ServiceGroupApp.JwtService
	roleService = service.ServiceGroupApp.SysRoleService
	menuService = service.ServiceGroupApp.MenuService
)

type ApiGroup struct {
	BaseApi
	JwtApi
	RoleApi
	MenuApi
}

var ApiGroupApp = new(ApiGroup)
