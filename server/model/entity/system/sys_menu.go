package system

import "github.com/zhengpanone/gin-vue3-admin/global"

type SysMenu struct {
	ID         string `gorm:"primarykey;->first" json:"ID"` // 主键ID
	Name       string `json:"name" gorm:"column:name;comment:菜单名称"`
	Icon       string `json:"icon" gorm:"column:icon;comment:菜单图标"`
	Path       string `json:"path" gorm:"column:path;comment:路由path"`
	ParentId   uint   `json:"parentId" gorm:"column:parent_id;comment:父菜单ID"`
	MenuLevel  uint   `json:"-" gorm:"column:menu_level;comment:菜单级别"`
	MenuType   uint   `json:"menuType" gorm:"column:type;comment:菜单类型,1:目录 2:菜单 3:按钮"`
	Visible    bool   `json:"visible" gorm:"column:visible;comment:是否在列表隐藏"`
	Permission string `json:"permission" gorm:"column:permission;comment:菜单权限标识"`
	Component  string `json:"component" gorm:"comment:前端文件路径"`
	Sort       uint   `json:"sort" gorm:"column:sort;comment:排序标记"`
	Meta       `json:"meta" gorm:"embedded;comment:附加属性"`
	Children   []SysMenu          `json:"children" gorm:"-"`
	Parameters []SysMenuParameter `json:"parameters" gorm:"-"`
	MenuBtn    []SysBaseMenuBtn   `json:"menuBtn" gorm:"-"`
	global.BaseModel
}

type Meta struct {
	ActiveName  string `json:"activeName" gorm:"comment:高亮菜单"`
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:是否缓存"`           // 是否缓存
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"` // 是否是基础路由（开发中）
	Title       string `json:"title" gorm:"comment:菜单名"`                // 菜单名
	Icon        string `json:"icon" gorm:"comment:菜单图标"`                // 菜单图标
	CloseTab    bool   `json:"closeTab" gorm:"comment:自动关闭tab"`         // 自动关闭tab
}

type SysMenuParameter struct {
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"comment:地址栏携带参数为params还是query"` // 地址栏携带参数为params还是query
	Key           string `json:"key" gorm:"comment:地址栏携带参数key"`             // 地址栏携带参数的key
	Value         string `json:"value" gorm:"comment:地址栏携带参数的值"`            // 地址栏携带参数的值
	global.BaseModel
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
