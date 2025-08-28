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
