package system

import (
	"github.com/zhengpanone/gin-vue3-admin/service"
)

type ApiGroup struct {
	BaseApi
	JwtApi
	RoleApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService  = service.ServiceGroupApp.SystemServiceGroup.JwtService
	roleService = service.ServiceGroupApp.SystemServiceGroup.SysRoleService
	menuService = service.ServiceGroupApp.SystemServiceGroup.MenuService
)
