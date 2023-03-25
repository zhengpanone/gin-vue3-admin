package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-vue3-admin/model/common/response"
)

// FindMenuByUserId
// @Tags FindMenuByUserId
// @Summary 根据用户ID查找该用户的菜单
func FindMenuByUserId(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	response.OkWithData(ctx, userId)
}
