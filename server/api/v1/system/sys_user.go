package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/common/request"
	"github.com/zhengpanone/gin-vue3-admin/model/common/response"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
	systemReq "github.com/zhengpanone/gin-vue3-admin/model/system/request"
	systemRes "github.com/zhengpanone/gin-vue3-admin/model/system/response"
	"github.com/zhengpanone/gin-vue3-admin/utils"

	"go.uber.org/zap"
)

// ChangePassword
//
//	@Tags			登陆注册
//	@Summary		更改密码
//	@Description	更改密码
//	@ID				/v1/user/changePassword
//	@Accept			json
//	@Produce		json
//	@Router			/v1/api/changePassword [post]
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
//
//	@Tags			登陆注册
//	@Summary		用户注册
//	@Description	用户注册
//	@ID				/v1/api/admin/register
//	@Accept			json
//	@Produce		application/json
//	@Param			data	body	request.RegisterParam	true	"body"	#[username,密码,手机号码]	body	[string,string,string]	[required,required,required]	"[system.SysUser]"
//	@Router			/v1/api/admin/register [post]
func (b *BaseApi) Register(ctx *gin.Context) {
	// 绑定参数
	var registerParam systemReq.RegisterParam
	_ = ctx.ShouldBindJSON(&registerParam)
	register, err := userService.Register(&registerParam)
	if err != nil {
		response.ErrorWithMsg(ctx, "注册失败"+err.Error())
		return
	}
	response.OkWithData(ctx, register)
}

// Logout
//
//	@Tags			登陆注册
//	@Summary		用户退出
//	@Description	退出登录
//	@ID				/v1/api/admin/logout
//	@Accept			json
//	@Produce		json
//	@Param			data	body	request.LoginParam	true	"用户名,密码,验证码,验证码ID"
//	@Router			/v1/api/admin/logout [post]
func (b *BaseApi) Logout(c *gin.Context) {

}

// Login
//
//		@Tags			登陆注册
//		@Summary		用户登录
//		@Description	用户登录
//		@ID				/v1/api/admin/login
//		@Accept			json
//		@Produce		application/json
//		@Param			data	body	request.LoginParam	true	"用户名,密码,验证码,验证码ID"
//	 	@Router			/v1/api/admin/login [post]
//		@Success  		200 {object} response.Response(data=systemRes.LoginResponse, msg=string) "返回包括用户信息，token过期时间"
func (b *BaseApi) Login(ctx *gin.Context) {
	var loginParam systemReq.LoginParam
	//key := ctx.ClientIP()
	err := ctx.ShouldBindJSON(&loginParam)
	if err != nil {
		response.ErrorWithMsg(ctx, "参数绑定失败")
		return
	}
	if loginParam.Username == "" || loginParam.Password == "" {
		response.ErrorWithMsg(ctx, "用户名和密码不能为空！")
		return
	}
	if global.GVA_CONFIG.Captcha.Enable == true {
		if !store.Verify(loginParam.CaptchaID, loginParam.Captcha, true) {
			response.ErrorWithMsg(ctx, "验证码校验失败")
			return
		}
	}
	user, err := userService.LoginPwd(&loginParam)
	if err != nil {
		global.GVA_LOG.Error("登录失败", zap.Any("user", loginParam))
		response.ErrorWithMsg(ctx, err.Error())
		return
	} else {
		b.tokenNext(ctx, *user)
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

// GetUserList
// @Tags      SysUser
// @Summary   分页获取用户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router    /user/getUserList [post]
func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	//err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	list, total, err := userService.GetUserInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.ErrorWithMsg(c, "获取失败")
		return
	}
	response.OkWithDataAndMsg(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功")
}

// SetUserInfo
// @Tags      SysUser
// @Summary   设置用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysUser                                             true  "ID, 用户名, 昵称, 头像链接"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "设置用户信息"
// @Router    /user/setUserInfo [put]
func (b *BaseApi) SetUserInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	// err = utils.Verify(user, utils.IdVerify)
	// if err != nil {
	// 	response.ErrorWithMsg(err.Error(), c)
	// 	return
	// }

}
