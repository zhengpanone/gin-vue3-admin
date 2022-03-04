package v1

import (
	"fmt"
	"gin-api-learn/global"
	"gin-api-learn/middleware"
	"gin-api-learn/model/entity"
	"gin-api-learn/model/request"
	"gin-api-learn/model/response"
	"github.com/zhengpanone/gin-api-learn/server/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	var loginParam request.LoginParam

	_ = ctx.ShouldBindJSON(&loginParam)
	fmt.Println("参数:", loginParam)
	if loginParam.Phone == "" || loginParam.Password == "" {
		response.Error(ctx, "手机号和密码不能为空！")
		return
	}
	// 调用登录服务
	userRecord := entity.User{Phone: loginParam.Phone, Password: loginParam.Password}
	if err := service.LoginPwd(&userRecord); err != nil {
		global.GlobalLogger.Error("登录失败", zap.Any("user", userRecord))
		response.Error(ctx, "登录失败,账号或密码错误")
		return
	}
	// 生成Token
	token, err := middleware.CreateToken(userRecord.ID)
	if err != nil {
		global.GlobalLogger.Sugar().Error("登录失败,Token生成异常：%s", err)
		response.Error(ctx, "登录失败,账号或密码错误")
	}
	userRecord.Token = token
	response.OkWithData(ctx, userRecord)
}

func GetUser(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.OkWithData(ctx, user)
}
