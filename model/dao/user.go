package dao

import (
	"gin-api-learn/global"
	"gin-api-learn/model/entity"
)

type UserDao struct {
	Uid uint
}

func (u *UserDao) FindUser() (*entity.User, error) {
	var user entity.User

	// 校验账号和密码
	result := global.GlobalMysqlClient.Where("id=? ", u.Uid).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	// 查询用户信息
	userInfo := entity.UserInfo{}
	result = global.GlobalMysqlClient.Where("uid = ? ", u.Uid).First(&userInfo)
	if result.Error != nil {
		return nil, result.Error
	}
	user.UserInfo = userInfo
	return &user, nil
}
