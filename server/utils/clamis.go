package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/request"
	"net"
)

func ClearToken(c *gin.Context) {
	//
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", "", -1, "/", "", false, false)
	} else {
		c.SetCookie("x-token", "", -1, "/", host, false, false)
	}
}

func GetClaims(c *gin.Context) (*request.CustomClaims, error) {
	token := c.Request.Header.Get("X-Token")
	if token == "" {
		token = c.GetHeader("Authorization")
	}
	if token == "" {
		return nil, errors.New("请求头中获取不到token")
	}

	j := NewJWT()
	claims, err := j.ParseToken(token)
	fmt.Printf("t4est%s", "userClaim")
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
			return cl.UUID
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.UUID
	}
}

func GetUserAuthorityId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.AuthorityId
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.AuthorityId
	}
}
