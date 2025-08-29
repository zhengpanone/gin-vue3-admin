package utils

import (
	"errors"
	systemReq "github.com/zhengpanone/gin-vue3-admin/request"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"go.uber.org/zap"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}

func GetToken(ctx *gin.Context) (string, error) {
	token, _ := ctx.Cookie("x-token")
	if token != "" {
		return token, nil
	}
	token = ctx.Request.Header.Get("x-token")
	if token != "" {
		return token, nil
	}
	token = ctx.Request.Header.Get("TOKEN")
	if token != "" {
		return token, nil
	}
	token = ctx.GetHeader("Authorization")
	if token != "" {
		return token, nil
	}
	//获取当前请求方法
	if ctx.Request.Method == http.MethodGet {
		token, ok := ctx.GetQuery("token")
		if ok {
			return token, nil
		}
	}
	if ctx.Request.Method == http.MethodPost {
		postParam := make(map[string]interface{})

		_ = ctx.ShouldBindJSON(&postParam)
		token, ok := postParam["token"]
		if ok {
			return token.(string), nil
		}
	}
	return "", errors.New("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在TOKEN且claims是否为规定结构")
}

func (j *JWT) CreateClaims(baseClaims systemReq.BaseClaims) systemReq.CustomClaims {
	claims := systemReq.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: global.GVA_CONFIG.JWT.BufferTime, //缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                             // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.GVA_CONFIG.JWT.ExpireTime, // 过期时间 7天  配置文件
			Issuer:    global.GVA_CONFIG.JWT.Issuer,                         // 签名发行者
			IssuedAt:  time.Now().Unix(),                                    // 签发时间
		},
	}
	return claims
}

// CreateToken 创建一个Token
func (j *JWT) CreateToken(claims systemReq.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.GVA_CONFIG.JWT.SigningKey))
}

// ParseToken 解析JWT
func (j *JWT) ParseToken(tokenString string) (*systemReq.CustomClaims, error) {
	tokenString = strings.Fields(tokenString)[0]
	token, err := jwt.ParseWithClaims(tokenString, &systemReq.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			global.GVA_LOG.Error("解析JWT失败", zap.String("err", err.Error()))
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}

		return nil, err
	}
	if token != nil {
		// 断言
		//验证
		if claims, ok := token.Claims.(*systemReq.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}

}
