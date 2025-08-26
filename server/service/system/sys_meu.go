package system

import (
	"errors"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
	"gorm.io/gorm"
)

type MenuService struct {
}

func (menuService *MenuService) AddBaseMenu(menu system.SysBaseMenu) error {
	if !errors.Is(global.GVA_DB.Where("name=?", menu.Name).First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name,请修改name")
	}
	return global.GVA_DB.Create(&menu).Error
}

func (menuService *MenuService) GetMenuTree(authorityId uint) (menus []system.SysMenu, err error) {
	menuTree, err := menuService.getMenuTreeMap(authorityId)
	menus = menuTree[0]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

func (menuService *MenuService) getMenuTreeMap(authorityId uint) (treeMap map[uint][]system.SysMenu, err error) {
	var allMenus []system.SysMenu
	var baseMenu []system.SysBaseMenu
	var btns []system.SysAuthorityBtn
	treeMap = make(map[uint][]system.SysMenu)

	var SysAuthorityMenus []system.SysAuthorityMenu
	err = global.GVA_DB.Where("sys_authority_authority_id=?", authorityId).Find(&SysAuthorityMenus).Error
	if err != nil {
		return
	}
	var MenuIds []string
	for i := range SysAuthorityMenus {
		MenuIds = append(MenuIds, SysAuthorityMenus[i].MenuId)
	}
	if len(MenuIds) == 0 {
		return nil, nil
	}
	err = global.GVA_DB.Where("id in (?)", MenuIds).Order("sort").Preload("Parameters").Find(&baseMenu).Error
	if err != nil {
		return
	}
	for i := range baseMenu {
		allMenus = append(allMenus, system.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthorityId: authorityId,
			MenuId:      baseMenu[i].ID,
			Parameters:  baseMenu[i].Parameters,
		})
	}
	err = global.GVA_DB.Where("authority_id=?", authorityId).Preload("SysBaseMenuBtn").Find(&btns).Error
	if err != nil {
		return
	}
	var btnMap = make(map[uint]map[string]uint)
	for _, v := range allMenus {
		v.Btns = btnMap[v.SysBaseMenu.ID]
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

// @function:		getChildrenList
// @description:	获取子菜单
// @param:			menu *model.SysMenu, treeMap map[string][]model.SysMenu
// @return:		err error
func (menuService *MenuService) getChildrenList(menu *system.SysMenu, treeMap map[uint][]system.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
