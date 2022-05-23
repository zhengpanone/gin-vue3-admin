package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/zhengpanone/gin-api-learn/server/config"
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
	GVA_CONFIG        config.ServerConfig
	GlobalMysqlClient *gorm.DB      // mysql客户端
	GlobalRedisClient *redis.Client // redis客户端
	GVA_LOG           *zap.Logger   // 日志
)
