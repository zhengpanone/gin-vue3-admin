package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/common/response"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
	"go.uber.org/zap"
)

type AuthorityMenuApi struct {
}

// AddBaseMenu
// @Tags Menu
// @Summary 新增菜单
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

// FindMenuByUserId
// @Tags    FindMenuByUserId
// @Summary 根据用户ID查找该用户的菜单
func FindMenuByUserId(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	response.OkWithData(ctx, userId)
}
