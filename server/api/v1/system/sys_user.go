package system

import (
	"fmt"
	systemReq "github.com/zhengpanone/gin-vue3-admin/model/system/request"

	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/common/response"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
	systemRes "github.com/zhengpanone/gin-vue3-admin/model/system/response"
	"github.com/zhengpanone/gin-vue3-admin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ChangePassword
// @Tags        SysUser
// @Summary     更改密码
// @Description 更改密码
// @ID          /v1/user/changePassword
// @Accept      json
// @Produce     json
// @Router      /v1/api/changePassword [post]
func (b *BaseApi) ChangePassword(ctx *gin.Context) {
	var changePassword systemReq.ChangePasswordParam
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

// Register
// @Tags        Base
// @Summary     用户注册
// @Description 用户注册
// @ID          /v1/api/admin/register
// @Accept      json
// @Produce     application/json
// @Param       data body request.RegisterParam true "body" #[username,密码,手机号码] body [string,string,string] [required,required,required] "[system.SysUser]"
// @Router      /v1/api/admin/register [post]
func (b *BaseApi) Register(ctx *gin.Context) {
	// 绑定参数
	var registerParam systemReq.RegisterParam
	_ = ctx.ShouldBindJSON(&registerParam)
	register, err := userService.Register(&registerParam)
	if err != nil {
		response.ErrorWithMsg(ctx, "注册失败")
		return
	}
	response.OkWithData(ctx, register)
}

// Logout
// @Tags        Base
// @Summary     用户退出
// @Description 退出登录
// @ID          /v1/api/admin/logout
// @Accept      json
// @Produce     json
// @Param       data body request.LoginParam true "用户名,密码,验证码,验证码ID"
// @Router      /v1/api/admin/logout [post]
func (b *BaseApi) Logout(c *gin.Context) {

}

// Login
// @Tags        Base
// @Summary     用户登录
// @Description 用户登录
// @ID          /v1/api/admin/login
// @Accept      json
// @Produce     json
// @Param       data body request.LoginParam true "用户名,密码,验证码,验证码ID"
// @Router      /v1/api/admin/login [post]
func (b *BaseApi) Login(ctx *gin.Context) {
	var loginParam systemReq.LoginParam
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
	claims := j.CreateClaims(systemReq.BaseClaims{
		Username: user.Username,
		UUID:     user.UUID,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Sugar().Error("登录失败,Token生成异常：%s", err)
		response.ErrorWithMsg(ctx, "获取token失败")
		return
	}
	user.Token = token
	var menus []systemRes.Menus
	menus = append(menus, systemRes.Menus{
		Path:   "首页",
		Title:  "首页",
		Header: "首页",
	})
	menus = append(menus, systemRes.Menus{
		Path:   "商品",
		Title:  "商品",
		Header: "商品",
		Children: []systemRes.Menus{
			{Title: "商品列表"},
			{Title: "商品分类"},
			{Title: "商品规格"},
			{Title: "商品评论"},
		},
	})
	response.OkWithDataAndMsg(ctx, systemRes.LoginResponse{
		User:      user,
		Token:     token,
		UserInfo:  systemRes.UserInfo{Id: user.UUID, Account: user.Username, HeadPic: ""},
		ExpiresAt: claims.ExpiresAt,
		Menus:     menus,
	}, "登录成功")
}

func (b *BaseApi) GetUser(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.OkWithData(ctx, user)
}
