package middleware

import (
	"github.com/gin-gonic/gin"
   )

// 跨域中间件

func CORS(ctx *gin.Context){
	method := ctx.Request.Method
	ctx.Header("Access-Control-Allow-Origin", ctx.Request.Header.Get("Origin"))
	ctx.Header("Access-Control-Allow-Credentials","true")
	ctx.Header("Access-Control-Allow-Headers","Content-Type,Access-Control-Allow-Headers,Authorization,X-Requested-With")
	ctx.Header("Access-Control-Allow-Methods","GET,POST,PUT,PATCH,DELETE,OPTIONS")

	if method=="OPTIONS" || method =="HEAD"{
		ctx.AborWithStatus(204)
		return 
	}
	ctx.Next()
}

// TODO 增加限流器 https://hongker.github.io/2021/02/18/golang-limiter/