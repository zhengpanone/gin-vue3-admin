package system

import "github.com/zhengpanone/gin-vue3-admin/global"

type SysBaseMenuBtn struct {
	Name          string `json:"name" gorm:"comment:按钮关键key"`
	Desc          string `json:"desc" gorm:"按钮备注"`
	SysBaseMenuID uint   `json:"sysBaseMenuID" gorm:"comment:菜单ID"`
	global.BaseModel
}
