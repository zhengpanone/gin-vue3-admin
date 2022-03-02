package main

import (
	"gin-api-learn/core"
	"gin-api-learn/initialize"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()
	core.RunServer()

}
