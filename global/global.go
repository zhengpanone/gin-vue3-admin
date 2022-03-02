package global

import (
	"gin-api-learn/config"

	"go.uber.org/zap"
)

/**
全局常量和变量
*/
const (
	CONFIGFILE = "../config/app.yml"
)

var (
	GlobalConfig config.ServerConfig
	GlobalLogger *zap.Logger // 日志
)
