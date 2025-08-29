package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/response"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

// Captcha
//
// @Tags		登陆注册
// @Summary	生成验证码
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Success		200	{object}	response.Response{data=response.SysCaptchaResponse,msg=string}	"生成验证码,返回包括随机数id,base64,验证码长度"
// @Router		/v1/api/captcha [get]
func (b *BaseApi) Captcha(c *gin.Context) {
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(
		global.GVA_CONFIG.Captcha.ImgHeight,
		global.GVA_CONFIG.Captcha.ImgWidth,
		global.GVA_CONFIG.Captcha.KeyLong,
		0.7,
		80)

	cp := base64Captcha.NewCaptcha(driver, store)

	if id, b64s, _, err := cp.Generate(); err != nil {
		global.GVA_LOG.Error("验证码获取失败", zap.Error(err))
		response.ErrorWithMsg(c, "验证码获取失败")
	} else {
		response.OkWithDataAndMsg(c, response.SysCaptchaResponse{
			CaptchaID:     id,
			PicPath:       b64s,
			CaptchaLength: global.GVA_CONFIG.Captcha.KeyLong,
		}, "验证码获取成功")
	}
}
