package dao

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
)

type UserDao struct {
	ID   int64
	UUID uuid.UUID
}

func (u *UserDao) FindUser() (*system.SysUser, error) {
	var user system.SysUser

	// 校验账号和密码
	result := global.GlobalMysqlClient.Where("id=? and uuid=? ", u.ID, u.UUID.String()).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
