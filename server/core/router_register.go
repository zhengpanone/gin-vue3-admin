package core

import (
	"gin-api-learn/router"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(engine *gin.Engine) {

	router.InitSystemRouter(engine)
	router.InitRedisRouter(engine)
	router.InitSwaggerRouter(engine)
	router.RouterGroupApp.System.InitUserRouter(engine)

}