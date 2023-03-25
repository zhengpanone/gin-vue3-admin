package system

import (
	"github.com/zhengpanone/gin-vue3-admin/global"
)

type SysPermissionPoint struct {
	global.BaseModel
	PointClass  string
	PointIcon   string
	PointStatus string
}
