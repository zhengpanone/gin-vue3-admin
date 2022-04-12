package system

import (
	"fmt"
	"gin-api-learn/global"
	"gin-api-learn/middleware"
	"gin-api-learn/model/request"
	"gin-api-learn/model/response"

	"gin-api-learn/service/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Register
// @Tags SysUser
// @Summary 用户注册
// @Description 用户注册
// @ID /v1/user/register
// @Accept  json
// @Produce  application/json
// @Param data body request.RegisterParam true "body" #[username,密码,手机号码] body [string,string,string] [required,required,required] "[system.SysUser]
// @Router /v1/user/register [post]
func Register(ctx *gin.Context) {
	// 绑定参数
	var registerParam request.RegisterParam
	_ = ctx.ShouldBindJSON(&registerParam)
	register, err := system.Register(registerParam)
	if err != nil {
		response.Error(ctx, "注册失败")
		return
	}
	response.OkWithData(ctx, register)
}

// ChangePassword
// @Tags SysUser
// @Summary 更改密码
// @Description 更改密码
// @ID /v1/user/changePassword
// @Accept  json
// @Produce  json
// @Router /v1/user/changePassword [post]
func ChangePassword(ctx *gin.Context) {
	var changePassword request.ChangePasswordStruct
	_ = ctx.ShouldBindJSON(&changePassword)
	if changePassword.Username == "" || changePassword.Password == "" || changePassword.NewPassword == "" {
		response.Error(ctx, "用户名、密码、新密码不能为空")
		return
	}
	err := system.ChangePassword(changePassword)
	fmt.Println(err)
}

// Login
// @Tags SysUser
// @Summary 登录
// @Description 用户登录
// @ID /v1/user/login
// @Accept json
// @Produce json
// @Router /v1/user/login [post]
func Login(ctx *gin.Context) {
	var loginParam request.LoginParam
	_ = ctx.ShouldBindJSON(&loginParam)
	if loginParam.Username == "" || loginParam.Password == "" {
		response.Error(ctx, "用户名和密码不能为空！")
		return
	}
	err, user := system.LoginPwd(&loginParam)
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

func GetUser(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.OkWithData(ctx, user)
}
