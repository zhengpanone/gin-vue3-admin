package router

import (
	"gin-api-learn/api/v1/system"
	"github.com/gin-gonic/gin"
)

func InitMenu(engine *gin.Engine) {
	menuRouter := engine.Group("/v1/menu")
	{
		menuRouter.GET("menu", system.FindMenuByUserId)
	}

}
