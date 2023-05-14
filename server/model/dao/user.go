package dao

import (
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
)

type UserDao struct {
	UserID string
}

func (u *UserDao) FindUser() (*system.SysUser, error) {
	var user system.SysUser

	// 校验账号和密码
	result := global.GlobalMysqlClient.Where("uid=? ", u.UserID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
