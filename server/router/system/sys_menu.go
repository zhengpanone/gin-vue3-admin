package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-vue3-admin/api/v1/system"
)

func InitMenu(engine *gin.Engine) {
	menuRouter := engine.Group("/v1/menu")
	{
		menuRouter.GET("menu", system.FindMenuByUserId)
	}

}
