package v1

import (
	"gin-api-learn/global"
	"gin-api-learn/model/response"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetConfig(ctx *gin.Context) {
	// Sugar模式
	global.GlobalLogger.Sugar().Infof("日志写入测试:%v", strings.Repeat("hello\t", 6))

	global.GlobalLogger.Info("Info记录", zap.String("name", "张三"))
	response.OkWithData(ctx, global.GlobalConfig)
}
