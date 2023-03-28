package system

import (
	"errors"
	"fmt"
	"github.com/zhengpanone/gin-vue3-admin/utils"

	uuid "github.com/satori/go.uuid"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
	"github.com/zhengpanone/gin-vue3-admin/model/request"

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

	var user system.SysUser
	err := global.GlobalMysqlClient.Transaction(func(tx *gorm.DB) error {
		if !errors.Is(tx.Where("username = ?", param.Username).First(&user).Error, gorm.ErrRecordNotFound) {
			return errors.New("用户名已经注册")
		}
		user = system.SysUser{
			Username: param.Username,
			Password: utils.MD5V([]byte(param.Password)),
			UUID:     uuid.NewV4(),
		}
		if err := tx.Create(&user).Error; err != nil {
			global.GVA_LOG.Sugar().Errorf("新增用户失败：%s", err)
			return err
		}
		userInfo := system.SysUserInfo{
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
	return &user, err
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
