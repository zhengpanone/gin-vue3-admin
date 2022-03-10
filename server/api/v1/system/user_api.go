package system

import (
	"fmt"
	"gin-api-learn/global"
	"gin-api-learn/middleware"
	"gin-api-learn/model/entity/system"
	"gin-api-learn/model/request"
	"gin-api-learn/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseApi struct{}

//@Tags SysUser
//@Summary 用户注册
// @Produce  application/json
// @Param data body systemReq.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {object} response.Response{data=systemRes.SysUserResponse,msg=string} "用户注册账号,返回包括用户信息"
// @Router /user/register [post]
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

func (b *BaseApi) Login(ctx *gin.Context) {
	var loginParam request.LoginParam

	_ = ctx.ShouldBindJSON(&loginParam)
	fmt.Println("参数:", loginParam)
	if loginParam.Phone == "" || loginParam.Password == "" {
		response.Error(ctx, "手机号和密码不能为空！")
		return
	}
	// 调用登录服务
	userRecord := system.SysUser{Phone: loginParam.Phone, Password: loginParam.Password}
	if err := userService.LoginPwd(&userRecord); err != nil {
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

func (b *BaseApi) GetUser(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.OkWithData(ctx, user)
}
