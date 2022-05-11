package utils

import (
	"gin-api-learn/global"
	"gin-api-learn/model/request"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func GetClaims(c *gin.Context) (*request.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GlobalLogger.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
		return nil, err
	}
	return claims, err

}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *gin.Context) (userID string) {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return uuid.UUID{}.String()
		} else {
			return cl.UserID
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.UserID
	}
}
