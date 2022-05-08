package core

import (
	"gin-api-learn/router"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(engine *gin.Engine) {
	//engine.Use(middleware.CORS())

	router.InitSystemRouter(engine)
	router.InitRedisRouter(engine)
	router.InitSwaggerRouter(engine)
	router.InitUserRouter(engine)
	router.InitMenu(engine)

}

func CORS() {
	panic("unimplemented")
}
