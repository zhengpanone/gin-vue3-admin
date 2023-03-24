package response

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const (
	SUCCESS      = http.StatusOK
	ERROR        = -1
	TOKEN_EXPIRE = -2
)

//Response  响应结构体
type Response struct {
	// 业务状态码
	Code int `json:"code"`
	// 提示信息
	Msg string `json:"msg"`
	// 响应数据
	Data interface{} `json:"data"`
	// Meta 源数据,存储如请求ID,分页等信息
	Meta Meta `json:"meta"`
	// Errors 错误提示，如 xx字段不能为空等
	Errors []ErrorItem `json:"errors"`
}

// Meta 元数据
type Meta struct {
	RequestId string `json:"request_id"`
	// TODO 还可以集成分页信息
}

// ErrorItem 错误项
type ErrorItem struct {
	Key   string `json:"key"`
	Value string `json:"error"`
}

func New() *Response {
	return &Response{
		Code:   200,
		Msg:    "",
		Data:   nil,
		Meta:   Meta{RequestId: uuid.NewV4().String()},
		Errors: []ErrorItem{},
	}
}

func ResultJson(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
		Meta: Meta{
			RequestId: strconv.FormatInt(time.Now().UnixNano(), 10),
		},
	})
	return
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

func OkWithDataAndMsg(ctx *gin.Context, data interface{}, msg string) {
	ResultJson(ctx, SUCCESS, msg, data)
}

//Error 错误信息
func ErrorWithMsg(ctx *gin.Context, msg string) {
	ResultJson(ctx, ERROR, msg, map[string]interface{}{})
}

func ErrorWithToken(ctx *gin.Context, msg string) {
	ResultJson(ctx, TOKEN_EXPIRE, msg, map[string]interface{}{})
}
