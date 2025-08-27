package utils

import "golang.org/x/crypto/bcrypt"

// BcryptPassword
//
//	@function:		BcryptPassword
//	@description:	bcrypt加密
//	@param:			string
//	@return:		string,error
func BcryptPassword(password string) (string, error) {
	// Cost 14 意味着 2^14 = 16,384 轮计算
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash
// 验证明文密码与哈希值
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
