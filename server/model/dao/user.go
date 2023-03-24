package dao

import (
	"github.com/zhengpanone/gin-api-learn/global"
	"github.com/zhengpanone/gin-api-learn/model/entity/system"
)

type UserDao struct {
	UserID string
}

func (u *UserDao) FindUser() (*system.SysUser, error) {
	var user system.SysUser

	// 校验账号和密码
	result := global.GlobalMysqlClient.Where("id=? ", u.UserID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	// 查询用户信息
	userInfo := system.SysUserInfo{}
	result = global.GlobalMysqlClient.Where("uid = ? ", u.UserID).First(&userInfo)
	if result.Error != nil {
		return nil, result.Error
	}
	user.UserInfo = userInfo
	return &user, nil
}
