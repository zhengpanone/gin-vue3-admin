package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/zhengpanone/gin-vue3-admin/docs"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/middleware"
	"github.com/zhengpanone/gin-vue3-admin/response"
	"github.com/zhengpanone/gin-vue3-admin/router"
)

// InitRouters 初始化总路由
func InitRouters() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp
	global.GVA_LOG.Info("use middleware logger")
	// http://127.0.0.1:8099/swagger/index.html
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	Router.Use(middleware.CORS())
	global.GVA_LOG.Info("use middleware cors")

	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			//response.WrapContext(c).Success(nil)
			response.Ok(c)
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由，不做鉴权
		systemRouter.InitRoleRouter(PublicGroup)
	}
	PrivateGroup := Router.Group("/v1/api")
	PrivateGroup.Use(middleware.JWTAuthMiddleware())
	{
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitJwtRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup) // 注册menu路由

	}

	global.GVA_LOG.Info("router register success")
	return Router
}
