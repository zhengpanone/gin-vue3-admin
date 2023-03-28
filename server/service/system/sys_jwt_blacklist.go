package system

import (
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
)

type JwtService struct{}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: JsonInBlacklist
// @description: 拉黑jwt
// @param: jwtList system.JwtBlacklist
// @return: err error
func (j *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.GlobalMysqlClient.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}
