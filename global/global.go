package global

import (
	"gin-api-learn/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

/**
全局常量和变量
*/
const (
	CONFIGFILE = "./config/app.yml"
)

var (
	GlobalConfig      config.ServerConfig
	GlobalMysqlClient *gorm.DB    // mysql客户端
	GlobalLogger      *zap.Logger // 日志
)
