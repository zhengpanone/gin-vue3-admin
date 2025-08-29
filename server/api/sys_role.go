package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model"
	"github.com/zhengpanone/gin-vue3-admin/request"
	sysRes "github.com/zhengpanone/gin-vue3-admin/response"
	"go.uber.org/zap"
)

type RoleApi struct{}

// CreateRole
//
//	@Tags			角色管理
//	@Summary		创建角色
//	@Description	创建角色
//	@Accept			json
//	@Produce		application/json
//	@Param			data	body	request.RoleParam	true	"body"	#[roleName,密码,手机号码]	body	[string,string,string]	[required,required,required]	"[system.SysUser]"
//	@Router			/v1/api/role/addRole [post]
func (r *RoleApi) CreateRole(c *gin.Context) {
	var roleParam request.RoleParam
	_ = c.ShouldBindJSON(&roleParam)
	role, err := roleService.CreateRole(roleParam)
	if err != nil {
		sysRes.ErrorWithMsg(c, err.Error())
		return
	}
	sysRes.OkWithData(c, role)
}

// PageRole
//
//	@Tags			角色管理
//	@Summary		分页获取角色
//	@Description	分页获取角色
//	@Accept			json
//	@Produce		application/json
//	@Param			data	body	request.PageRequest	true	"body"	#[roleName,密码,手机号码]	body	[string,string,string]	[required,required,required]	"[system.SysUser]"
//	@Router			/v1/api/role/pageRole [post]
func (r *RoleApi) PageRole(c *gin.Context) {
	var pageInfo request.PageRequest
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		sysRes.ErrorWithMsg(c, err.Error())
		return
	}
	list, total, err := roleService.PageRole(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败！", zap.Error(err))
		sysRes.ErrorWithMsg(c, err.Error())
		return
	}
	sysRes.OkWithDataAndMsg(c, sysRes.PageResult[model.SysRole]{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "")
}

// DeleteRole
//
//	@Tags			角色管理
//	@Summary		删除角色
//	@Description	删除角色
//	@Accept			json
//	@Produce		application/json
//	@Param			id	query		string				true	"角色ID"
//	@Success		200	{object}	response.Response	返回结果	200	类型（object就是结构体）	类型	注释
//	@Router			/v1/api/role/delete/:id [delete]
func (r *RoleApi) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	err := roleService.DeleteRole(id)
	if err != nil {
		sysRes.ErrorWithMsg(c, err.Error())
		return
	}
	sysRes.OkWithData(c, "删除成功")
}

// UpdateRole
//
//	@Tags			角色管理
//	@Summary		更新角色
//	@Description	update role
//	@Accept			json
//	@Produce		application/json
//	@Param			data	body	request.PageInfo	true	"body"	#[roleName,密码,手机号码]	body	[string,string,string]	[required,required,required]	"[system.SysUser]"
//	@Router			/v1/api/role/update [put]
func (r *RoleApi) UpdateRole(c *gin.Context) {
	id := c.Param("id")
	sysRes.OkWithData(c, id)
}
