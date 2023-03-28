package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/zhengpanone/gin-vue3-admin/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/*
*
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
	BlackCache        local_cache.Cache
)
