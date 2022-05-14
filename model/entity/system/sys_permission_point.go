package system

import "gin-api-learn/global"

type SysPermissionPoint struct {
	global.BaseModel
	PointClass  string
	PointIcon   string
	PointStatus string
}
