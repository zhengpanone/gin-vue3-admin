package response

import (
	"github.com/zhengpanone/gin-vue3-admin/model"
)

type SysMenusResponse struct {
	Menus []model.SysMenu `json:"menus"`
}
