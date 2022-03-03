package initialize

import (
	"fmt"
	"gin-api-learn/global"
)

func InitConfig() {
	InitViperConfig()
	InitLogger()
	InitGorm()
	InitRedis()
}

func CloseResource() {
	// TODO 关闭资源
	if global.GlobalMysqlClient != nil {
		db, _ := global.GlobalMysqlClient.DB()
		_ = db.Close()
	}

	if global.GlobalRedisClient != nil {
		_ = global.GlobalRedisClient.Close()
		fmt.Println("redis close")
	}
}
