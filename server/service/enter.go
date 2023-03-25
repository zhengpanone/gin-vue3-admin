package service

import (
	"github.com/zhengpanone/gin-vue3-admin/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
