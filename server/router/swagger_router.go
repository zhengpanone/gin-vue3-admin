package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitSwaggerRouter(engine *gin.Engine) {
	redisRouter := engine.Group("swagger")
	{
		redisRouter.GET("*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
