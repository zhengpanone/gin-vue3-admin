package initialize

import (
	"fmt"

	"github.com/zhengpanone/gin-vue3-admin/global"
)

func InitConfig() {
	InitViperConfig() //初始化Viper
	InitLogger()      // 初始化log
	InitGorm()
	// if global.GlobalRedisClient
	// InitRedis()
}

func CloseResource() {
	//  关闭资源
	if global.GVA_DB != nil {
		db, _ := global.GVA_DB.DB()
		_ = db.Close()
	}

	if global.GlobalRedisClient != nil {
		_ = global.GlobalRedisClient.Close()
		fmt.Println("redis close")
	}
}
