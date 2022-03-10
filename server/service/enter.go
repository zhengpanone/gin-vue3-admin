package service

import "gin-api-learn/service/system"

type ServiceGroup struct {
	SystemServiceGroup system.SysUserService
}

var ServiceGroupApp = new(ServiceGroup)
