package initialize

import (
	"gin-api-learn/global"
	"gin-api-learn/router"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化总路由
func InitRouters() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System
	global.GVA_LOG.Info("use middleware logger")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	//engine.Use(middleware.CORS())

	// router.InitSystemRouter(engine)
	// router.InitRedisRouter(engine)
	// router.InitSwaggerRouter(engine)
	// router.InitUserRouter(engine)
	// router.InitMenu(engine)
	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由，不做鉴权
	}
	global.GVA_LOG.Info("router register success")
	return Router
}

func CORS() {
	panic("unimplemented")
}
