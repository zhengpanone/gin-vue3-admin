package core

import (
	"gin-api-learn/router"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(engine *gin.Engine) {
	router.InitSystemRouter(engine)
	router.InitUserRouter(engine)
	router.InitRedisRouter(engine)
}
