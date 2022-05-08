package middleware

import (
	"errors"
	"gin-api-learn/global"
	"gin-api-learn/model/dao"
	"gin-api-learn/model/request"
	"gin-api-learn/model/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

// JWT中间件
func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 获取参数中的token
		token := getToken(ctx)
		global.GlobalLogger.Sugar().Info("token:%s", token)
		if token == "" {
			response.Error(ctx, "Token不能为空")
			// 中断请求
			ctx.Abort()
			return
		}
		// 验证Token
		userClaim, err := ParseToken(token)
		if err != nil {
			response.ErrorWithToken(ctx, "Token error:"+err.Error())
			ctx.Abort()
			return
		}
		// 设置到上下文
		setContextData(ctx, userClaim, token)
		// 继续请求后续流程
		ctx.Next()

	}
}

// 设置数据到上下文
func setContextData(ctx *gin.Context, userClaim *request.UserClaims, token string) {
	userDao := &dao.UserDao{
		UserID: userClaim.UserID,
	}
	user, err := userDao.FindUser()
	if err != nil {
		response.Error(ctx, "用户不存在")
		ctx.Abort()
		return
	}
	user.Token = token
	ctx.Set("userClaim", userClaim)
	ctx.Set("user", user)
}

func getToken(ctx *gin.Context) string {
	token := ctx.Request.Header.Get("TOKEN")
	if token != "" {
		return token
	}
	token = ctx.GetHeader("Authorization")
	if token != "" {
		return token
	}
	//获取当前请求方法
	if ctx.Request.Method == http.MethodGet {
		token, ok := ctx.GetQuery("token")
		if ok {
			return token
		}
	}
	if ctx.Request.Method == http.MethodPost {
		postParam := make(map[string]interface{})

		_ = ctx.ShouldBindJSON(&postParam)
		token, ok := postParam["token"]
		if ok {
			return token.(string)
		}
	}
	return ""
}

// 创建JWT
func CreateToken(userID string) (string, error) {
	newWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		&request.UserClaims{
			StandardClaims: &jwt.StandardClaims{
				ExpiresAt: time.Now().Add(global.GlobalConfig.Jwt.Expire).Unix(),
				Issuer:    global.GlobalConfig.Jwt.Issuer, // 签发人
				IssuedAt:  time.Now().Unix(),              // 签发时间
			},
			UserID: userID,
		})
	return newWithClaims.SignedString([]byte(global.GlobalConfig.Jwt.Secret))
}

// 解析JWT
func ParseToken(tokenString string) (*request.UserClaims, error) {
	var err error
	var token *jwt.Token
	token, err = jwt.ParseWithClaims(tokenString, &request.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.GlobalConfig.Jwt.Secret), nil
	})
	if err != nil {
		global.GlobalLogger.Error("解析JWT失败", zap.String("err", err.Error()))
		return nil, err
	}
	// 断言
	userClaims, ok := token.Claims.(*request.UserClaims)
	// 验证
	if !ok || !token.Valid {
		return nil, errors.New("JWT校验失败")
	}
	return userClaims, nil
}
