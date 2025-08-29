package model

import "github.com/zhengpanone/gin-vue3-admin/global"

type JwtBlacklist struct {
	global.BaseModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
