package system

type SysMenu struct {
	SysBaseMenu
	MenuId      uint                   `json:"menu_id" gorm:"comment:菜单ID"`
	AuthorityId uint                   `json:"-" gorm:"comment:角色ID"`
	Children    []SysMenu              `json:"children" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"perm" gorm:"comment:foreignKey:SysBaseMenuID;references:MenuId"`
	Btns        map[string]uint        `json:"btns" gorm:"-"`
}

type SysAuthorityMenu struct {
	MenuId      string `json:"menu_id" gorm:"comment:菜单ID;column:sys_base_menu_id"`
	AuthorityId string `json:"-" gorm:"comment:角色ID;column:sys_authority_authority_id"`
}

func (s SysAuthorityMenu) TableName() string {
	return "sys_authority_menus"
}
