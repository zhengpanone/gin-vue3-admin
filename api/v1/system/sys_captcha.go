package system

import (
	"gin-api-learn/global"
	"gin-api-learn/model/response"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

// Captcha
// @Tags Base
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Sucess 200 {object}
// @Router /base/captha [post]
func (b *BaseApi) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.GVA_LOG.Error("验证码获取失败", zap.Error(err))
		response.ErrorWithMsg(c, "验证码获取失败")
	} else {
		response.OkWithData(c)
	}
}
