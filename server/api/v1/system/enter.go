package system

import (
	"github.com/zhengpanone/gin-api-learn/server/service"
)

type ApiGroup struct {
	BaseApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
