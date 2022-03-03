package router

import (
	v1 "gin-api-learn/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRedisRouter(engine *gin.Engine) {
	redisRouter := engine.Group("redis")
	{
		redisRouter.GET("test", v1.RedisPing)
	}
}
