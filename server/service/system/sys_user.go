package system

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/common/request"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
	systemReq "github.com/zhengpanone/gin-vue3-admin/model/system/request"
	"github.com/zhengpanone/gin-vue3-admin/utils"
	"gorm.io/gorm"
)

type UserService struct {
}

// LoginPwd
// @author: [zhengpanone](https://github.com/zhengpanone)
// @function: LoginPwd
// @description: 用户登录
// @param: u *request.LoginParam
// @return: userInfo *system.SysUser, err error
func (userService *UserService) LoginPwd(u *systemReq.LoginParam) (userInfo *system.SysUser, err error) {
	var user system.SysUser
	err = global.GVA_DB.Where("username=?", u.Username).First(&user).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	checkResult := utils.CheckPasswordHash(u.Password, user.Password)
	if !checkResult {
		return nil, errors.New("密码错误")
	}
	return &user, err
}

//@author: [zhengpanone](https://github.com/zhengpanone)
//@function: Register
//@description: 用户注册
//@param: param *request.RegisterParam
//@return: user *system.SysUser,err error

func (userService *UserService) Register(param *systemReq.RegisterParam) (*system.SysUser, error) {

	var user system.SysUser
	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if !errors.Is(tx.Where("username = ?", param.Username).First(&user).Error, gorm.ErrRecordNotFound) {
			return errors.New("用户名已经注册")
		}
		bcryptPassword, _ := utils.BcryptPassword(param.Password)
		user = system.SysUser{
			Username: param.Username,
			Password: bcryptPassword,
			UUID:     uuid.Must(uuid.NewV4()),
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
func (userService *UserService) ChangePassword(param systemReq.ChangePasswordParam) error {
	result := global.GVA_DB.Where("username=? and password=?", param.Username, param.Password).First(param)
	fmt.Println(result)
	return result.Error
}

// GetUserInfo 获取用户信息
// TODO
func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user system.SysUser, err error) {
	var sysUser system.SysUser
	//err = global.GVA_DB.Preload("")
	return sysUser, err
}

//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (userService *UserService) GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}
