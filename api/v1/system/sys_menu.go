package system

import (
	"gin-api-learn/model/response"
	"github.com/gin-gonic/gin"
)

// FindMenuByUserId
// @Tags FindMenuByUserId
// @Summary 根据用户ID查找该用户的菜单
func FindMenuByUserId(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	response.OkWithData(ctx, userId)
}
