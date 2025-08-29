package service

import (
	"github.com/zhengpanone/gin-vue3-admin/dao"
	"github.com/zhengpanone/gin-vue3-admin/model"
	"github.com/zhengpanone/gin-vue3-admin/request"
	"github.com/zhengpanone/gin-vue3-admin/response"
)

type MenuService struct {
}

func (menuService *MenuService) AddMenu(dto request.AddMenu) error {
	menu := model.SysMenu{
		Name:       dto.Name,
		MenuType:   dto.MenuType,
		Permission: dto.Permission,
		Sort:       dto.Sort,
	}
	menuDao := dao.SysMenuDao{}
	if err := menuDao.ValidateMenuName(menu.ParentId, menu.Name); err != nil {
		return err
	}
	_, err := menuDao.AddMenu(menu)
	return err
}

func (menuService *MenuService) UpdateMenu(menu model.SysMenu) error {
	menuDao := dao.SysMenuDao{}
	return menuDao.UpdateMenu(menu)
}

func (menuService *MenuService) DeleteMenuByIds(ids []interface{}) error {
	menuDao := dao.SysMenuDao{}
	return menuDao.DeleteMenuByIds(ids)
}

func (menuService *MenuService) GetMenuById(id interface{}) (menu model.SysMenu, err error) {
	menuDao := dao.SysMenuDao{}
	menu, err = menuDao.GetMenuById(id)
	return
}

func (menuService *MenuService) GetMenuList(queryMenu request.QueryMenu) (result response.PageResult[model.SysMenu], err error) {
	menuDao := dao.SysMenuDao{}
	result, err = menuDao.PageMenu(queryMenu)
	return result, err
}
