package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/common/request"
	"github.com/zhengpanone/gin-vue3-admin/model/common/response"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
	systemRes "github.com/zhengpanone/gin-vue3-admin/model/system/response"
	"github.com/zhengpanone/gin-vue3-admin/utils"
	"go.uber.org/zap"
)

type AuthorityMenuApi struct {
}

// AddBaseMenu
// @Tags      Menu
// @Summary   新增菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysBaseMenu             true  "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success   200   {object}  response.Response{msg=string}  "新增菜单"
// @Router    /v1/api/menu/addBaseMenu [post]
func (a *AuthorityMenuApi) AddBaseMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	err = menuService.AddBaseMenu(menu)
	if err != nil {
		global.GVA_LOG.Error("添加失败", zap.Error(err))
		response.ErrorWithMsg(c, "添加失败")
		return
	}
	response.OkWithMsg(c, "添加成功")
}

// GetMenu godoc
// @Tags AuthorityMenu
// @Summary   获取用户动态路由
// @Description  根据用户ID获取菜单
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      request.Empty  true  "空"
// @Success   200   {object}  response.Response{data=response.SysMenusResponse,msg=string}  "获取用户动态路由,返回包括系统菜单详情列表"
// @Router    /v1/api/menu/getMenu [post]
func (a *AuthorityMenuApi) GetMenu(c *gin.Context) {
	menus, err := menuService.GetMenuTree(utils.GetUserAuthorityId(c))
	if err != nil {
		global.GVA_LOG.Error("获取菜单失败！", zap.Error(err))
		response.ErrorWithMsg(c, "获取菜单失败")
		return
	}
	if menus == nil {
		menus = []system.SysMenu{}
	}
	response.OkWithDataAndMsg(c, systemRes.SysMenusResponse{Menus: menus}, "获取菜单成功")
}

// GetMenuList
// @Tags      Menu
// @Summary   分页获取基础menu列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取基础menu列表,返回包括列表,总数,页码,每页数量"
// @Router    /v1/api/menu/getMenuList [post]
func (a *AuthorityMenuApi) GetMenuList(c *gin.Context) {
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
