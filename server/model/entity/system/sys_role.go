package system

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zhengpanone/gin-vue3-admin/global"
)

// SysRole 角色表
type SysRole struct {
	global.BaseModel
	RoleID   uuid.UUID `json:"uuid" gorm:"not null;unique;primary_key;comment:角色ID;"` // 角色ID
	RoleName string    `json:"roleName" gorm:"comment:角色名称"`                          // 角色名
	RoleCode string    `json:"code" gorm:"comment:角色代码"`
	ParentId uuid.UUID `json:"parentId" gorm:"父角色ID"` // 父角色ID
	Remark   string    `json:"remark" gorm:"comment:备注"`
}

// SysUserRole 用户-角色
type SysUserRole struct {
	global.BaseModel
	UserID  string  `json:"userID" gorm:"comment:用户ID"`
	SysUser SysUser `gorm:"foreignkey:UserID"`
	RoleID  string  `json:"roleID" gorm:"comment:角色ID"`
	SysRole SysRole `gorm:"foreignkey:RoleID"`
}

// SysRolePermission 角色-权限关系表
type SysRolePermission struct {
	RoleID string `json:"roleId" gorm:"comment:角色ID"`
	MenuID string `json:"permission" gorm:"comment:"`
}
