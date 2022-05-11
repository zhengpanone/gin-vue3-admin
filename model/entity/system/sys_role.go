package system

import "gin-api-learn/global"

// SysRole 角色表
type SysRole struct {
	Name   string         `json:"name" gorm:"comment:角色名称"`
	Code   string         `json:"code" gorm:"comment:角色代码"`
	Remark string         `json:"remark" gorm:"comment:备注"`
	User   *[]SysUserRole `json:"sysUserRoleList" gorm:"foreignKey:RoleID;comment:角色用户列表"`
	Menu   *[]SysRoleMenu `json:"sysRoleMenu" gorm:"foreignKey:RoleID;"`
	global.BaseModel
}

// SysUserRole 用户-角色
type SysUserRole struct {
	UserID string `json:"userID" gorm:"not null;comment:用户ID"`
	RoleID string `json:"roleID" gorm:"not null;comment:角色ID"`
	global.BaseModel
}

//SysRoleMenu 角色-菜单关系表
type SysRoleMenu struct {
	RoleID string `json:"roleId" gorm:"not null;comment:角色ID"`
	MenuID string `json:"menuID" gorm:"not null;comment:菜单ID"`
	global.BaseModel
}
