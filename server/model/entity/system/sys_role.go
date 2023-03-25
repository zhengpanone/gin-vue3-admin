package system

import (
	"github.com/zhengpanone/gin-vue3-admin/global"
)

// SysRole 角色表
type SysRole struct {
	global.BaseModel
	Name   string `json:"name" gorm:"comment:角色名称"`
	Code   string `json:"code" gorm:"comment:角色代码"`
	Remark string `json:"remark" gorm:"comment:备注"`
}

// SysUserRole 用户-角色
type SysUserRole struct {
	global.BaseModel
	UserID  string  `json:"userID" gorm:"comment:用户ID"`
	SysUser SysUser `gorm:"foreignkey:UserID"`
	RoleID  string  `json:"roleID" gorm:"comment:角色ID"`
	SysRole SysRole `gorm:"foreignkey:RoleID"`
}

// SysRoleMenu 角色-菜单关系表
type SysRoleMenu struct {
	RoleID string `json:"roleId" gorm:"comment:角色ID"`
	MenuID string `json:"menuID" gorm:"comment:"`
}
