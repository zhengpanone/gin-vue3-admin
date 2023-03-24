package system

import (
	"github.com/zhengpanone/gin-api-learn/global"
	"github.com/zhengpanone/gin-api-learn/model/common/response"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetConfig
// @Tags Base
// @Summary 获取配置
// @accept application/json
// @Produce application/json
// @Router /v1/api/config [get]
func (b *BaseApi) GetConfig(ctx *gin.Context) {
	// Sugar模式
	global.GVA_LOG.Sugar().Infof("日志写入测试:%v", strings.Repeat("hello\t", 6))

	global.GVA_LOG.Info("Info记录", zap.String("name", "张三"))
	response.OkWithData(ctx, global.GVA_CONFIG)
}
