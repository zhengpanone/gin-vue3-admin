package request

import (
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	UUID     uuid.UUID
	UserID   int64 // UserId
	Username string
	NickName string
	RoleId   string
}
