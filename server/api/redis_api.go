package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/response"
)

func RedisPing(ctx *gin.Context) {
	method, _ := ctx.GetQuery("type")
	var result string
	var err error
	switch method {
	case "get":

		result, err = global.GlobalRedisClient.Get(ctx, "test").Result()
	case "set":
		err = global.GlobalRedisClient.Set(ctx, "test", "hello world", 0).Err()
	}

	if err != nil {
		fmt.Println("------------")
		fmt.Println(err)
		response.ErrorWithMsg(ctx, err.Error())
		return
	}

	response.OkWithData(ctx, result)

}
