package system

import (
	"errors"
	"fmt"
	"gin-api-learn/global"
	"gin-api-learn/model/entity/system"
	"gin-api-learn/model/request"
	uuid "github.com/satori/go.uuid"

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
	if err == nil {
		// TODO 查询目录
	}
	return &user, err
}

//@author: [zhengpanone](https://github.com/zhengpanone)
//@function: Register
//@description: 用户注册
//@param: param *request.RegisterParam
//@return: user *system.SysUser,err error

func (userService *UserService) Register(param *request.RegisterParam) (*system.SysUser, error) {
	user := &system.SysUser{
		UserName: param.Username,
		Phone:    param.Phone,
		Password: param.Password,
		RoleID:   param.RoleID,
	}
	user, err := findUserByUserName(user)
	user.ID = uuid.NewV4().String()
	if err != nil {
		return user, errors.New("用户已注册")
	}
	err = global.GlobalMysqlClient.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			global.GlobalLogger.Sugar().Errorf("新增用户失败：%s", err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, nil
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

func findUserByUserName(sysUser *system.SysUser) (*system.SysUser, error) {

	result := global.GlobalMysqlClient.Where("user_name=?", sysUser.UserName).First(&sysUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return sysUser, result.Error
	}
	return sysUser, nil
}
