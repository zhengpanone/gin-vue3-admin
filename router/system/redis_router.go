package system

import (
	"gin-api-learn/api/v1/system"

	"github.com/gin-gonic/gin"
)

func InitRedisRouter(engine *gin.Engine) {
	redisRouter := engine.Group("redis")
	{
		redisRouter.GET("test", system.RedisPing)
	}
}
