package system

import (
	"fmt"
	"gin-api-learn/global"
	"gin-api-learn/model/entity/system"
	"gin-api-learn/model/request"

	"gorm.io/gorm"
)

func LoginPwd(u *request.LoginParam) (err error, userInfo *system.SysUser) {
	var user system.SysUser
	err = global.GlobalMysqlClient.Where("username=? and password=?", u.Username, u.Password).First(&user).Error
	if err == nil {
		// TODO 查询目录
	}
	return err, &user
}

func Register(param request.RegisterParam) (*system.SysUser, error) {
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

// 更改用户名
func ChangePassword(param request.ChangePasswordStruct) error {
	result := global.GlobalMysqlClient.Where("username=? and password=?", param.Username, param.Password).First(param)
	fmt.Println(result)
	return result.Error
}
