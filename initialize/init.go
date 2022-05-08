package initialize

import (
	"fmt"
	"gin-api-learn/global"
)

func InitConfig() {
	InitViperConfig() //初始化Viper
	InitLogger()      // 初始化log
	InitGorm()
	//if global.GlobalRedisClient
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
