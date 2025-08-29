package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model"
	"github.com/zhengpanone/gin-vue3-admin/request"
	"github.com/zhengpanone/gin-vue3-admin/response"
	"go.uber.org/zap"
)

type MenuApi struct {
}

// AddMenu
//
// @Tags		菜单管理
// @Summary	新增菜单
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		request.AddMenu	true	"菜单数据"
// @Success		200		{object}	response.Response{msg=string}	"更新菜单"
// @Router		/v1/api/menu/addMenu [post]
func (a *MenuApi) AddMenu(c *gin.Context) {
	var addMenu request.AddMenu
	err := c.ShouldBindJSON(&addMenu)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}

	err = menuService.AddMenu(addMenu)
	if err != nil {
		global.GVA_LOG.Error("添加失败", zap.Error(err))
		response.ErrorWithMsg(c, err.Error())
		return
	}
	response.OkWithMsg(c, "添加成功")
}

// UpdateMenu
//
// @Tags		菜单管理
// @Summary		更新菜单
// @Security	ApiKeyAuth
// @accept 		application/json
// @Produce 	application/json
// @Param       id    	path      		string             				true  "菜单ID"
// @Param       data  	body      		request.UpdateMenu   		true  "菜单数据"
// @Success		200		{object}		response.Response{msg=string}	"更新菜单"
// @Router		/v1/api/menu/updateMenu [put]
func (a *MenuApi) UpdateMenu(c *gin.Context) {
	id := c.Param("id")
	var updateMenu request.UpdateMenu
	err := c.ShouldBindJSON(&updateMenu)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	menu := model.SysMenu{
		ID:         id,
		Name:       updateMenu.Name,
		MenuType:   updateMenu.MenuType,
		Permission: updateMenu.Permission,
		Sort:       updateMenu.Sort,
	}
	err = menuService.UpdateMenu(menu)
	if err != nil {
		global.GVA_LOG.Error("更新失败", zap.Error(err))
		response.ErrorWithMsg(c, err.Error())
		return
	}
	response.OkWithMsg(c, "更新成功")
}

// DeleteMenuByIds
//
// @Tags		菜单管理
// @Summary	删除菜单
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param       data   	body      	request.CommonIds  true  "菜单ID数组"
// @Success		200		{object}	response.Response{msg=string}	"更新菜单"
// @Router		/v1/api/menu/deleteMenu [delete]
func (a *MenuApi) DeleteMenuByIds(c *gin.Context) {
	var idsObj request.CommonIds
	err := c.ShouldBindJSON(&idsObj)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}

	err = menuService.DeleteMenuByIds(idsObj.Ids)
	if err != nil {
		global.GVA_LOG.Error("删除失败", zap.Error(err))
		response.ErrorWithMsg(c, err.Error())
		return
	}
	response.OkWithMsg(c, "删除成功")
}

// GetMenuById
// @Tags		菜单管理
// @Summary		获取菜单
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param       id		query     	string     true  "菜单ID"
// @Success		200		{object}	response.Response{msg=string}	"更新菜单"
// @Router		/v1/api/menu/getMenuById [get]
func (a *MenuApi) GetMenuById(c *gin.Context) {
	id := c.Query("id")
	menu, err := menuService.GetMenuById(id)
	if err != nil {
		global.GVA_LOG.Error("获取菜单详情失败", zap.Error(err))
		response.ErrorWithMsg(c, err.Error())
		return
	}
	response.OkWithData(c, menu)
}

// GetMenuList
//
// @Tags		Menu
// @Summary		分页获取基础menu列表
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		request.PageRequest										true
// @Success		200		{object}	response.Response{data=response.PageResult,msg=string}	"分页获取基础menu列表,返回包括列表,总数,页码,每页数量"
// @Router		/v1/api/menu/getMenuList [post]
func (a *MenuApi) GetMenuList(c *gin.Context) {
	var queryMenu request.QueryMenu
	err := c.ShouldBindJSON(&queryMenu)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	result, err := menuService.GetMenuList(queryMenu)

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.ErrorWithMsg(c, "获取失败")
		return
	}
	response.OkWithDataAndMsg(c, result, "获取成功")
}
