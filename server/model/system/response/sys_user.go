package response

import (
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
)

type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
	Menus     []Menus        `json:"menus"`
}

type Menus struct {
	Path     string  `json:"path"`
	Title    string  `json:"title"`
	Icon     string  `json:"icon"`
	Header   string  `json:"header"`
	IsHeader int32   `json:"isHeader"`
	Children []Menus `json:"children"`
}
