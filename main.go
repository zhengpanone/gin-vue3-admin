package main

import (
	"gin-api-learn/core"
	"gin-api-learn/initialize"
)

/**
 * @Author: zhengpanone
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2022/04/07 上午10:11
 */
func main() {
	initialize.InitConfig()
	core.RunServer()

}
