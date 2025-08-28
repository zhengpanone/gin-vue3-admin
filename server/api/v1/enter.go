package v1

import (
	"github.com/zhengpanone/gin-vue3-admin/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	system.MenuApi
}

var ApiGroupApp = new(ApiGroup)
