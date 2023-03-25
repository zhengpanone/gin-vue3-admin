package system

import (
	"github.com/zhengpanone/gin-vue3-admin/global"
)

type SysPermission struct {
	global.BaseModel
	Name        string `json:"name" gorm:"comment:权限名称"`
	Type        int8   `json:"type" gorm:"comment:权限类型：1为菜单 2为功能 3 为API"`
	Code        string `json:"code" gorm:"comment:权限编码"`
	Description string `json:"description" gorm:"权限描述"`
	PId         string `json:"pid" gorm:"comment:父ID"`
	EnVisible   int8   `json:"enVisible" gorm:"comment:是否可见"`
}
