package dao

import (
	"errors"

	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
	"gorm.io/gorm"
)

type SysRoleDao struct{}

// @author: [zhengpanone](https://github.com/zhengpanone)
// @function: AddRole
// @description: 添加角色
// @param: role system.SysRole
// @return: r system.SysRole, err error
func (sr *SysRoleDao) AddRole(r system.SysRole) (system.SysRole, error) {
	var role system.SysRole
	if !errors.Is(global.GlobalMysqlClient.Where("role_name=?", r.RoleName).First(&role).Error, gorm.ErrRecordNotFound) {
		return r, errors.New("角色已存在")
	}
	err := global.GlobalMysqlClient.Create(&r).Error
	return r, err
}

// @author: [zhengpanone](https://github.com/zhengpanone)
// @function: FindRole
// @description: 查找角色
// @param: role system.SysRole
// @return: r system.SysRole, err error
func (sr *SysRoleDao) FindRole(role system.SysRole) system.SysRole {
	var r system.SysRole
	global.GlobalMysqlClient.Where("role_name=?", role.RoleName).First(&r)
	return r
}
