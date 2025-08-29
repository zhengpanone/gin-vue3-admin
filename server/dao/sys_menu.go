package dao

import (
	"errors"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model"
	"github.com/zhengpanone/gin-vue3-admin/request"
	"github.com/zhengpanone/gin-vue3-admin/response"
	"github.com/zhengpanone/gin-vue3-admin/utils"
	"gorm.io/gorm"
)

type SysMenuDao struct{}

// ValidateMenuName 校验参数名称是否重复
func (sm *SysMenuDao) ValidateMenuName(parentId uint, menuName string) error {
	if !errors.Is(global.GVA_DB.Where("parent_id=? and name=?", parentId, menuName).First(&model.SysMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复名称,请修改名称")
	}
	return nil
}

// AddMenu 添加菜单
func (sm *SysMenuDao) AddMenu(r model.SysMenu) (model.SysMenu, error) {
	err := global.GVA_DB.Create(&r).Error
	return r, err
}

// UpdateMenu 修改菜单
func (sm *SysMenuDao) UpdateMenu(menu model.SysMenu) error {
	err := global.GVA_DB.Model(&model.SysMenu{}).Where("id=?", menu.ID).Updates(&menu).Error
	return err
}

// DeleteMenuByIds 根据Ids删除菜单
func (sm *SysMenuDao) DeleteMenuByIds(Ids []interface{}) error {
	err := global.GVA_DB.Delete(&model.SysMenu{}, "id in ?", Ids).Error
	return err
}

// GetMenuById 根据Id获取菜单详情
func (sm *SysMenuDao) GetMenuById(Id interface{}) (menu model.SysMenu, err error) {
	err = global.GVA_DB.Where("id = ?", Id).First(&menu).Error
	return
}

// PageMenu 根据Id获取菜单详情
func (sm *SysMenuDao) PageMenu(queryMenu request.QueryMenu) (menu response.PageResult[model.SysMenu], err error) {
	result, err := utils.Paginate[model.SysMenu](global.GVA_DB, queryMenu.PageRequest)
	return result, err
}
