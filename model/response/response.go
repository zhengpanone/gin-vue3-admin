package response

import (
	"gin-api-learn/global"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS      = 0
	ERROR        = -1
	TOKEN_EXPIRE = -2
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time string      `json:"-"`
}

func ResultJson(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
		Time: time.Now().Format(global.YYYYMMDDHHIISS),
	})
}

func Ok(ctx *gin.Context) {
	ResultJson(ctx, SUCCESS, "请求成功", map[string]interface{}{})
}

func OkWithMsg(ctx *gin.Context, msg string) {
	ResultJson(ctx, SUCCESS, msg, nil)
}

func OkWithData(ctx *gin.Context, data interface{}) {
	ResultJson(ctx, SUCCESS, "请求成功", data)
}
