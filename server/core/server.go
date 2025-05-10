package core

import (
	"fmt"

	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/initialize"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 获取自定义HTTP SERVER
func getCustomHttpServer(engine *gin.Engine) *http.Server {
	// 创建自定义配置服务
	httpServer := &http.Server{
		//ip和端口号
		Addr: global.GVA_CONFIG.App.Addr,
		//调用的处理器，如为nil会调用http.DefaultServeMux
		Handler: engine,
		//计算从成功建立连接到request body(或header)完全被读取的时间
		ReadTimeout: time.Second * 10,
		//计算从request body(或header)读取结束到 response write结束的时间
		WriteTimeout: time.Second * 10,
		//请求头的最大长度，如为0则用DefaultMaxHeaderBytes
		MaxHeaderBytes: 1 << 20,
	}
	return httpServer
}

// RunServer 启动服务
func RunServer() {
	Router := initialize.InitRouters()
	httpServer := getCustomHttpServer(Router)

	printServerInfo()
	_ = httpServer.ListenAndServe()
}

func printServerInfo() {
	appConfig := global.GVA_CONFIG.App
	// http://127.0.0.1:8099/swagger/index.html
	fmt.Printf("\n【 当前环境: %s 当前版本: %s 接口地址: http://%s 】\n", appConfig.Env, appConfig.Version, appConfig.Addr)
}
