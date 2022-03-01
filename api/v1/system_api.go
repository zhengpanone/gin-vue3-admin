package v1

import (
	"gin-api-learn/global"
	"gin-api-learn/model/response"

	"github.com/gin-gonic/gin"
)

func GetConfig(ctx *gin.Context) {
	response.OkWithData(ctx, global.GlobalConfig)
}
