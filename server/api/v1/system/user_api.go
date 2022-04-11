package system

import (
	"fmt"
	"gin-api-learn/global"
	"gin-api-learn/middleware"
	"gin-api-learn/model/request"
	"gin-api-learn/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseApi struct{}

func (b *BaseApi) Register(ctx *gin.Context) {
	// 绑定参数
	var registerParam request.RegisterParam
	_ = ctx.ShouldBindJSON(&registerParam)
	register, err := userService.Register(registerParam)
	if err != nil {
		response.Error(ctx, "注册失败")
		return
	}
	response.OkWithData(ctx, register)
}

// ChangePassword godoc
// @Summary 管理员登陆
// @Description 管理员登陆
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param body body request.ChangePasswordStruct true "body"    # [值得名称] body [值得类型] [是否必传] "[返回值名称]"
func (b *BaseApi) ChangePassword(ctx *gin.Context) {
	var changePassword request.ChangePasswordStruct
	_ = ctx.ShouldBindJSON(&changePassword)
	if changePassword.Username == "" || changePassword.Password == "" || changePassword.NewPassword == "" {
		response.Error(ctx, "用户名、密码、新密码不能为空")
		return
	}
	err := userService.ChangePassword(changePassword)
	fmt.Println(err)
}

func (b *BaseApi) Login(ctx *gin.Context) {
	var loginParam request.LoginParam
	_ = ctx.ShouldBindJSON(&loginParam)
	if loginParam.Username == "" || loginParam.Password == "" {
		response.Error(ctx, "用户名和密码不能为空！")
		return
	}
	err, user := userService.LoginPwd(&loginParam)
	if err != nil {
		global.GlobalLogger.Error("登录失败", zap.Any("user", loginParam))
		response.Error(ctx, "登录失败,账号或密码错误")
		return
	}
	// 生成Token
	token, err := middleware.CreateToken(user.ID)
	if err != nil {
		global.GlobalLogger.Sugar().Error("登录失败,Token生成异常：%s", err)
		response.Error(ctx, "登录失败,账号或密码错误")
	}
	user.Token = token
	response.OkWithData(ctx, user)
}

func (b *BaseApi) GetUser(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.OkWithData(ctx, user)
}
