package router

import (
	"github.com/zhengpanone/gin-vue3-admin/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
