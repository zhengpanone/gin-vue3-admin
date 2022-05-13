package initialize

import (
	"context"
	"gin-api-learn/global"

	"github.com/go-redis/redis/v8"
)

func InitRedis() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     global.GVA_CONFIG.Redis.Addr,
		Password: global.GVA_CONFIG.Redis.Password,
		DB:       global.GVA_CONFIG.Redis.DefaultDB,
	})
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), global.GVA_CONFIG.Redis.DialTimeout)
	defer cancelFunc()
	_, err := redisClient.Ping(timeoutCtx).Result()
	if err != nil {
		panic("redis初始化失败! " + err.Error())
	}
	global.GlobalRedisClient = redisClient
}
