package response

import "github.com/zhengpanone/gin-vue3-admin/model/entity/system"

type SysMenusResponse struct {
	Menus []system.SysMenu `json:"menus"`
}
