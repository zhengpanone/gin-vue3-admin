package main

import (
	"gin-api-learn/core"
	"gin-api-learn/initialize"
)

// @title gin learn API （必填，缺少会有警告）
// @version 1.0 （必填）
// @description This is gin learn api docs.
// @license.name Apache 2.0
// @contact.name go-swagger帮助文档
// @contact.url https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @host localhost:8080
// @BasePath /api/v1
func main() {
	initialize.InitConfig()
	core.RunServer()

}
