package system

import (
	"gin-api-learn/global"
	"gin-api-learn/model/entity/system"
	"gin-api-learn/model/request"

	"gorm.io/gorm"
)

type SysUserService struct{}

func (userService *SysUserService) LoginPwd(user *system.SysUser) error {
	result := global.GlobalMysqlClient.Where("phone=? and password=?", user.Phone, user.Password).First(user)
	return result.Error
}

func (userService *SysUserService) Register(param request.RegisterParam) (*system.SysUser, error) {
	user := system.SysUser{
		Username: param.Username,
		Phone:    param.Phone,
		Password: param.Password,
	}
	global.GlobalMysqlClient.Transaction(func(tx *gorm.DB) error {
		// tx.Where("username = ?", user.Username).First()

		if err := tx.Create(&user).Error; err != nil {
			global.GlobalLogger.Sugar().Errorf("新增用户失败：%s", err)
			return err
		}
		userInfo := system.SysUserInfo{
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
