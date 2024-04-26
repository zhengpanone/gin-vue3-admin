package request

import (
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	UserID   uuid.UUID // UserId
	Username string
	NickName string
	RoleId   string
}
