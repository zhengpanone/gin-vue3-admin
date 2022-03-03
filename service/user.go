package service

import (
	"gin-api-learn/global"
	"gin-api-learn/model/entity"
	"gin-api-learn/model/request"

	"gorm.io/gorm"
)

func LoginPwd(user *entity.User) error {
	result := global.GlobalMysqlClient.Where("phone=? and password=?", user.Phone, user.Password).First(user)
	return result.Error
}

func Register(param request.RegisterParam) (*entity.User, error) {
	user := entity.User{
		NickName: param.NickName,
		Phone:    param.Phone,
		Password: param.Password,
	}
	global.GlobalMysqlClient.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			global.GlobalLogger.Sugar().Errorf("新增用户失败：%s", err)
			return err
		}
		userInfo := entity.UserInfo{
			Uid:      user.ID,
			Birthday: param.Birthday,
			Address:  param.Address,
		}
		if err := tx.Create(&userInfo).Error; err != nil {
			global.GlobalLogger.Sugar().Errorf("新增用户信息失败：%s", err)
			return err
		}
		user.UserInfo = userInfo
		return nil
	})
	return &user, nil
}
