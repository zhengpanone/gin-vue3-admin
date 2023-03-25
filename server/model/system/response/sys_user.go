package response

import "github.com/zhengpanone/gin-vue3-admin/model/entity/system"

type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
	Menus     string         `json:"menus"`
}
