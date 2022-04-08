package core

import (
	"gin-api-learn/router"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(engine *gin.Engine) {
	router.Use(CORS)
	router.InitSystemRouter(engine)
	router.InitRedisRouter(engine)
}
