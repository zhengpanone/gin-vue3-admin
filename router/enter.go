package router

import (
	"github.com/zhengpanone/gin-api-learn/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
