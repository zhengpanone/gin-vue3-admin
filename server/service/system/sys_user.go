package system

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/zhengpanone/gin-api-learn/global"
	"github.com/zhengpanone/gin-api-learn/model/entity/system"
	"github.com/zhengpanone/gin-api-learn/model/request"

	"gorm.io/gorm"
)

type UserService struct {
}

//@author: [zhengpanone](https://github.com/zhengpanone)
//@function: LoginPwd
//@description: 用户登录
//@param: u *request.LoginParam
//@return: userInfo *system.SysUser, err error

func (userService *UserService) LoginPwd(u *request.LoginParam) (userInfo *system.SysUser, err error) {
	var user system.SysUser
	err = global.GlobalMysqlClient.Where("username=? and password=?", u.Username, u.Password).First(&user).Error
	if err != nil {
		return nil, errors.New("密码错误")
	}
	return &user, err
}

//@author: [zhengpanone](https://github.com/zhengpanone)
//@function: Register
//@description: 用户注册
//@param: param *request.RegisterParam
//@return: user *system.SysUser,err error

func (userService *UserService) Register(param *request.RegisterParam) (*system.SysUser, error) {
	user := system.SysUser{
		Username: param.Username,
		Phone:    param.Phone,
		Password: param.Password,
	}
	global.GlobalMysqlClient.Transaction(func(tx *gorm.DB) error {
		// tx.Where("username = ?", user.Username).First()

		if err := tx.Create(&user).Error; err != nil {
			global.GVA_LOG.Sugar().Errorf("新增用户失败：%s", err)
			return err
		}
		userInfo := system.SysUserInfo{
			UserID:   user.ID,
			Birthday: param.Birthday,
			Address:  param.Address,
		}
		if err := tx.Create(&userInfo).Error; err != nil {
			global.GVA_LOG.Sugar().Errorf("新增用户信息失败：%s", err)
			return err
		}
		user.UserInfo = userInfo
		return nil
	})
	return &user, nil
}

// ChangePassword 更改密码
func (userService *UserService) ChangePassword(param request.ChangePasswordParam) error {
	result := global.GlobalMysqlClient.Where("username=? and password=?", param.Username, param.Password).First(param)
	fmt.Println(result)
	return result.Error
}

// GetUserInfo 获取用户信息
// TODO
func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user system.SysUser, err error) {
	var sysUser system.SysUser
	//err = global.GlobalMysqlClient.Preload("")
	return sysUser, err
}
