package dao

import (
	"errors"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
	"gorm.io/gorm"
)

type SysMenuDao struct{}

// ValidateMenuName 校验参数名称是否重复
func (sm *SysMenuDao) ValidateMenuName(parentId uint, menuName string) error {
	if !errors.Is(global.GVA_DB.Where("parent_id=? and name=?", parentId, menuName).First(&system.SysMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复名称,请修改名称")
	}
	return nil
}

// AddMenu 添加菜单
func (sm *SysMenuDao) AddMenu(r system.SysMenu) (system.SysMenu, error) {
	err := global.GVA_DB.Create(&r).Error
	return r, err
}
