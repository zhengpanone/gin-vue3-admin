package model

// SysRole 角色表
type SysRole struct {
	ID         uint      `gorm:"primarykey;->first" json:"ID"`                    // 主键ID
	RoleCode   string    `json:"roleCode" gorm:"column:role_code;comment:角色Code"` // 角色Code
	RoleName   string    `json:"roleName" gorm:"column:role_name;comment:角色名称"`   // 角色名
	RoleDesc   string    `json:"roleDesc" gorm:"column:role_desc;comment:备注"`
	ParentCode string    `json:"parentCode" gorm:"column:parent_code;comment:父角色Code"` // 父角色Code
	Children   []SysRole `json:"children" gorm:"-"`
	BaseModel
}

// SysUserRole 用户-角色
type SysUserRole struct {
	UserID  string  `json:"userID" gorm:"column:user_id;comment:用户ID"`
	SysUser SysUser `gorm:"foreignkey:UserID"`
	RoleID  string  `json:"roleID" gorm:"column:role_id;comment:角色ID"`
	SysRole SysRole `gorm:"foreignkey:RoleID"`
	BaseModel
}

// SysRolePermission 角色-权限关系表
type SysRolePermission struct {
	RoleID string `json:"roleId" gorm:"comment:角色ID"`
	MenuID string `json:"permission" gorm:"comment:"`
	BaseModel
}

func (SysRole) TableName() string {
	return "sys_role"
}
