package main

import (
	"github.com/zhengpanone/gin-vue3-admin/core"
	"github.com/zhengpanone/gin-vue3-admin/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title			gin swagger API接口文档
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @version		0.0.1
// @description	This is gin learn api docs.
// @license.name	Apache 2.0
// @contact.name	go-swagger帮助文档
// @contact.url	https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @BasePath		/
func main() {
	initialize.InitConfig()
	core.RunServer()

}
