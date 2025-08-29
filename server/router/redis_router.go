package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-vue3-admin/api"
)

func InitRedisRouter(engine *gin.Engine) {
	redisRouter := engine.Group("redis")
	{
		redisRouter.GET("test", api.RedisPing)
	}
}
