package v1

import (
	"github.com/zhengpanone/gin-api-learn/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
