package core

import (
	"gin-api-learn/middleware"
	"gin-api-learn/router"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(engine *gin.Engine) {
	engine.Use(middleware.CORS())

	router.InitSystemRouter(engine)
	router.InitRedisRouter(engine)
}

func CORS() {
	panic("unimplemented")
}
