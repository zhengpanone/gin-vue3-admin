package dao

import (
	"github.com/gofrs/uuid/v5"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model"
)

type UserDao struct {
	UserID uuid.UUID
}

func (u *UserDao) FindUser() (*model.SysUser, error) {
	var user model.SysUser

	// 校验账号和密码
	result := global.GVA_DB.Where("uuid=? ", u.UserID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
