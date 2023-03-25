package main

import (
	"github.com/zhengpanone/gin-vue3-admin/core"
	"github.com/zhengpanone/gin-vue3-admin/initialize"
)

// @title gin api learn
// @version 0.0.1
// @description This is gin learn api docs.
// @license.name Apache 2.0
// @contact.name go-swagger帮助文档
// @contact.url https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @BasePath /
func main() {
	initialize.InitConfig()
	core.RunServer()

}
