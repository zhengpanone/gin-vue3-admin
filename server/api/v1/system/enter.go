package system

import (
	"github.com/zhengpanone/gin-vue3-admin/service"
)

type ApiGroup struct {
	BaseApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
