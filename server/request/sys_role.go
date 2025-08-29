package request

type RoleParam struct {
	RoleCode   string `json:"roleCode" binding:"required"`                 // 角色code
	RoleName   string `json:"roleName" binding:"required"`                 // 角色名
	ParentCode string `json:"parentCode"`                                  // 父角色Code
	RoleDesc   string `json:"roleDesc" gorm:"column:role_desc;comment:备注"` // 备注
}
