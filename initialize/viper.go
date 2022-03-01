package initialize

import (
	"flag"
	"fmt"
	"gin-api-learn/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// viper 初始化配置解析,函数可接受命令行参数
func InitConfig() {
	var configFile string
	flag.StringVar(&configFile, "c", global.CONFIGFILE, "配置文件路径")
	if len(configFile) == 0 {
		panic("配置文件不存在!")
	}
	// 读取配置文件
	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("配置文件解析失败:%s\n", err))
	}
	// 动态监测配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生改变")
		if err := v.Unmarshal(&global.GlobalConfig); err != nil {
			panic(fmt.Errorf("配置重置失败:%s\n", err))
		}
	})
	if err := v.Unmarshal(&global.GlobalConfig); err != nil {
		panic(fmt.Errorf("配置重载失败:%s\n", err))
	}
	// 设置配置文件
	global.GlobalConfig.App.ConfigFile = configFile
}
