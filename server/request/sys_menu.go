package request

type AddMenu struct {
	Name       string `json:"name" binding:"required"`     // 菜单名称
	MenuType   uint   `json:"menuType" binding:"required"` // 菜单类型,1:目录 2:菜单 3:按钮
	Permission string `json:"permission" `                 // 菜单权限标识
	Sort       uint   `json:"sort" `                       // 菜单排序
}

type UpdateMenu struct {
	ID         uint   `json:"id" binding:"required"`       // 菜单ID
	Name       string `json:"name" binding:"required"`     // 菜单名称
	Icon       string `json:"icon"`                        // 菜单icon
	MenuType   uint   `json:"menuType" binding:"required"` // 菜单类型,1:目录 2:菜单 3:按钮
	Permission string `json:"permission" `                 // 菜单权限标识
	Sort       uint   `json:"sort" `                       // 菜单排序
}

type QueryMenu struct {
	PageRequest PageRequest // 分页参数
	Name        string      `json:"name" binding:"required"`     // 菜单名称
	Icon        string      `json:"icon"`                        // 菜单icon
	MenuType    uint        `json:"menuType" binding:"required"` // 菜单类型,1:目录 2:菜单 3:按钮
	Permission  string      `json:"permission" `                 // 菜单权限标识
	Sort        uint        `json:"sort" `                       // 菜单排序

}
