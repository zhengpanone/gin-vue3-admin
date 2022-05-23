package router

import (
	"github.com/zhengpanone/gin-api-learn/server/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
