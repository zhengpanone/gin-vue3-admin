package router

import (
	v1 "gin-api-learn/api/v1"

	"github.com/gin-gonic/gin"
)

// 查看系统配置路由
func InitSystemRouter(engine *gin.Engine) {
	systemRouter := engine.Group("system")
	{
		systemRouter.GET("config", v1.GetConfig)
	}
}
