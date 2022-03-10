package system

import (
	"fmt"
	"gin-api-learn/global"
	"gin-api-learn/model/response"

	"github.com/gin-gonic/gin"
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
		response.Error(ctx, err.Error())
		return
	}

	response.OkWithData(ctx, result)

}
