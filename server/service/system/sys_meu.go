package system

import (
	"github.com/zhengpanone/gin-vue3-admin/model/dao"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
)

type MenuService struct {
}

func (menuService *MenuService) AddMenu(menu system.SysMenu) error {
	menuDao := dao.SysMenuDao{}
	if err := menuDao.ValidateMenuName(menu.ParentId, menu.Name); err != nil {
		return err
	}
	_, err := menuDao.AddMenu(menu)
	return err
}

func (menuService *MenuService) UpdateMenu(menu system.SysMenu) error {
	menuDao := dao.SysMenuDao{}
	return menuDao.UpdateMenu(menu)
}

func (menuService *MenuService) DeleteMenuByIds(ids []interface{}) error {
	menuDao := dao.SysMenuDao{}
	return menuDao.DeleteMenuByIds(ids)
}

func (menuService *MenuService) GetMenuById(id interface{}) (menu system.SysMenu, err error) {
	menuDao := dao.SysMenuDao{}
	menu, err = menuDao.GetMenuById(id)
	return
}
