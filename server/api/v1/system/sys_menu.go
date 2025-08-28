package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/common/request"
	"github.com/zhengpanone/gin-vue3-admin/model/common/response"
	sysEntity "github.com/zhengpanone/gin-vue3-admin/model/entity/system"
	sysRequest "github.com/zhengpanone/gin-vue3-admin/model/system/request"
	"go.uber.org/zap"
)

type MenuApi struct {
}

// AddMenu
//
//	@Tags		菜单管理
//	@Summary	新增菜单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		sysRequest.AddMenu				true	"路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
//	@Success	200		{object}	response.Response{msg=string}	"新增菜单"
//	@Router		/v1/api/menu/addMenu [post]
func (a *MenuApi) AddMenu(c *gin.Context) {
	var addMenu sysRequest.AddMenu
	err := c.ShouldBindJSON(&addMenu)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	menu := sysEntity.SysMenu{
		Name:       addMenu.Name,
		MenuType:   addMenu.MenuType,
		Permission: addMenu.Permission,
		Sort:       addMenu.Sort,
	}
	err = menuService.AddMenu(menu)
	if err != nil {
		global.GVA_LOG.Error("添加失败", zap.Error(err))
		response.ErrorWithMsg(c, err.Error())
		return
	}
	response.OkWithMsg(c, "添加成功")
}

// GetMenuList
//
//	@Tags		Menu
//	@Summary	分页获取基础menu列表
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		request.PageInfo										true	"页码, 每页大小"
//	@Success	200		{object}	response.Response{data=response.PageResult,msg=string}	"分页获取基础menu列表,返回包括列表,总数,页码,每页数量"
//	@Router		/v1/api/menu/getMenuList [post]
func (a *MenuApi) GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	/*err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}*/
	/*menuList, total, err := menuService.GetInfoList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}*/
	response.OkWithDataAndMsg(c, response.PageResult{
		List:     nil,
		Total:    1,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功")
}
