package test

import (
	"fmt"
	"testing"

	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/utils"
)

func TestCreateToken(t *testing.T) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)}
	/*claims := j.CreateClaims(request.BaseClaims{
		Username: "admin",
	})*/

	/*token, err := j.CreateToken(claims)
	if err != nil {
		fmt.Println(err)
	}*/
	// userClaim, err := j.ParseToken(token)
	userClaim, err := j.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIiLCJVc2VybmFtZSI6ImFkbWluIiwiTmlja05hbWUiOiIiLCJSb2xlSWQiOiIiLCJCdWZmZXJUaW1lIjowLCJleHAiOjE2ODAwMTQ1MTcsImlhdCI6MTY4MDAxNDUxNywiaXNzIjoiemhlbmdwYW5vbmUiLCJuYmYiOjE2ODAwMTM1MTd9.50qldZCQ_Cjn0nPWK_TEmeOISScyED7LNr9erS7C9Oc")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userClaim)
}
