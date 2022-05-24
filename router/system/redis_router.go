package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-api-learn/api/v1/system"
)

func InitRedisRouter(engine *gin.Engine) {
	redisRouter := engine.Group("redis")
	{
		redisRouter.GET("test", system.RedisPing)
	}
}
