package system

import "gin-api-learn/global"

type SysBaseMenuBtn struct {
	global.BaseModel
	Name          string `json:"name" gorm:"comment:按钮名称"`
	Desc          string `json:"desc" gorm:"comment:按钮描述"`
	SysBaseMenuID uint   `json:"sysBaseMenuID" gorm:"comment:菜单ID"`
}
