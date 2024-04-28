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
	UUID        uuid.UUID // UserId
	ID          uint
	Username    string
	NickName    string
	AuthorityId uint
}
