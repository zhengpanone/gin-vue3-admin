package model

import "github.com/zhengpanone/gin-vue3-admin/global"

type SysPermissionApi struct {
	global.BaseModel
	ApiUrl    string `json:"apiUrl" gorm:"comment:链接"`
	ApiMethod int8   `json:"apiMethod" gorm:"comment: 请求类型"`
	ApiLevel  string `json:"apiLevel" gorm:"comment:权限等级: 1为通用接口权限 2为需要校验接口权限"`
}
