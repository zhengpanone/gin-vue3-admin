package system

import (
	"github.com/zhengpanone/gin-api-learn/global"
)

type SysPermissionPoint struct {
	global.BaseModel
	PointClass  string
	PointIcon   string
	PointStatus string
}
