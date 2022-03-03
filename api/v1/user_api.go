package v1

import (
	"gin-api-learn/model/request"
	"gin-api-learn/model/response"
	"gin-api-learn/service"

	"github.com/gin-gonic/gin"
)

// 用户注册
func Register(ctx *gin.Context) {
	// 绑定参数
	var registerParam request.RegisterParam
	_ = ctx.ShouldBindJSON(&registerParam)
	// TODO 参数校验
	register, err := service.Register(registerParam)
	if err != nil {
		response.Error(ctx, "注册失败")
		return
	}
	response.OkWithData(ctx, register)
}

func Login(ctx *gin.Context) {

}
