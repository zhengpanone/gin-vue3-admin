package global

import "gin-api-learn/config"

/**
全局常量和变量
*/
const (
	CONFIGFILE = "../config/app.yml"
)

var (
	GlobalConfig config.ServerConfig
)
