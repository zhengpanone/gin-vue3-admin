package system

import "gin-api-learn/global"

type SysBaseMenu struct {
	global.BaseModel
	MenuLevel uint             `json:"-"`
	ParentId  string           `json:"parentId" gorm:"comment:父菜单ID"` // 父菜单ID
	Path      string           `json:"path" gorm:"comment:路由path"`    // 路由path
	Name      string           `json:"name" gorm:"comment:路由name"`    // 路由name
	Children  []SysBaseMenu    `json:"children" gorm:"-"`
	MenuBtn   []SysBaseMenuBtn `json:"menuBtn"`
}
