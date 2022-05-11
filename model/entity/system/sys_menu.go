package system

import "gin-api-learn/global"

type SysMenu struct {
	ParentId string         `json:"parentId" gorm:"comment:父菜单ID,一级菜单为0"`
	Name     string         `json:"name" gorm:"comment:菜单名称"`
	Path     string         `json:"path" gorm:"菜单URL"`
	Perms    string         `json:"perm" gorm:"comment:授权(例: sys:user:list)"`
	Type     string         `json:"type" gorm:"comment:类型, 0:目录; 1:菜单;2:按钮"`
	Icon     string         `json:"icon" gorm:"comment:菜单图标"`
	OrderNum uint           `json:"orderNum" gorm:"comment:排序"`
	Role     *[]SysRoleMenu `json:"sysRoleMenu"gorm:"foreignKey:MenuID;""`
	global.BaseModel
}
