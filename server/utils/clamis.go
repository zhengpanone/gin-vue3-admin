package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"github.com/zhengpanone/gin-vue3-admin/global"
	systemReq "github.com/zhengpanone/gin-vue3-admin/model/system/request"
)

func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
		return nil, err
	}
	return claims, err

}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *gin.Context) (userID uuid.UUID) {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UserID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.UserID
	}
}
