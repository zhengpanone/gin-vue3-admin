package system

import (
	"fmt"

	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/common/response"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
	"github.com/zhengpanone/gin-vue3-admin/model/request"
	systemRes "github.com/zhengpanone/gin-vue3-admin/model/system/response"
	"github.com/zhengpanone/gin-vue3-admin/utils"

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
// @Param data body request.RegisterParam true "body" #[username,密码,手机号码] body [string,string,string] [required,required,required] "[system.SysUser]"
// @Router /v1/user/register [post]
func Register(ctx *gin.Context) {
	// 绑定参数
	var registerParam request.RegisterParam
	_ = ctx.ShouldBindJSON(&registerParam)
	register, err := userService.Register(&registerParam)
	if err != nil {
		response.ErrorWithMsg(ctx, "注册失败")
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
// @Router /v1/api/changePassword [post]
func ChangePassword(ctx *gin.Context) {
	var changePassword request.ChangePasswordParam
	_ = ctx.ShouldBindJSON(&changePassword)
	if changePassword.Username == "" || changePassword.Password == "" || changePassword.NewPassword == "" {
		response.ErrorWithMsg(ctx, "用户名、密码、新密码不能为空")
		return
	}
	err := userService.ChangePassword(changePassword)
	fmt.Println(err)
}

func (b *BaseApi) GetUserInfo(ctx *gin.Context) {
	// var userParam request.LoginParam
	// _ = ctx.ShouldBindJSON(&userParam)
	// err := userService.GetUserInfo(userParam)
	// if err != nil {
	// 	response.Error(ctx, "获取用户信息失败")
	// 	return
	// }

}

// Login
// @Tags Base
// @Summary 用户登录
// @Description 用户登录
// @ID /v1/user/login
// @Accept json
// @Produce json
// @Param data body request.LoginParam true "用户名,密码,验证码,验证码ID"
// @Router /v1/api/login [post]
func (b *BaseApi) Login(ctx *gin.Context) {
	var loginParam request.LoginParam
	_ = ctx.ShouldBindJSON(&loginParam)
	if loginParam.Username == "" || loginParam.Password == "" {
		response.ErrorWithMsg(ctx, "用户名和密码不能为空！")
		return
	}
	if store.Verify(loginParam.CaptchaID, loginParam.Captcha, true) {
		user, err := userService.LoginPwd(&loginParam)
		if err != nil {
			global.GVA_LOG.Error("登录失败", zap.Any("user", loginParam))
			response.ErrorWithMsg(ctx, "登录失败,账号或密码错误")
			return
		} else {
			b.tokenNext(ctx, *user)
		}

	} else {
		response.ErrorWithMsg(ctx, "验证码校验失败")
	}

}

// 登录校验成功后签发jwt
func (b *BaseApi) tokenNext(ctx *gin.Context, user system.SysUser) {
	// 生成Token
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)}
	claims := j.CreateClaims(request.BaseClaims{
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Sugar().Error("登录失败,Token生成异常：%s", err)
		response.ErrorWithMsg(ctx, "获取token失败")
	}
	user.Token = token
	response.OkWithDataAndMsg(ctx, systemRes.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.ExpiresAt,
	}, "登录成功")
}

func GetUser(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.OkWithData(ctx, user)
}
