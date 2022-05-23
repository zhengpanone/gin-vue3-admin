package system

import (
	_ "github.com/zhengpanone/gin-api-learn/server/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSwaggerRouter(engine *gin.Engine) {
	redisRouter := engine.Group("swagger")
	{ // 文档界面访问URL
		// http://127.0.0.1:8099/swagger/index.html
		redisRouter.GET("*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
