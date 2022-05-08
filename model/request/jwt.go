package request

import "github.com/golang-jwt/jwt"

type UserClaims struct {
	*jwt.StandardClaims
	UserID string // 不建议放入会被修改的字段，推荐放入用户ID。
}
