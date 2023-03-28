package test

import (
	"fmt"
	"testing"

	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/request"
	"github.com/zhengpanone/gin-vue3-admin/utils"
)

func TestCreateToken(t *testing.T) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)}
	claims := j.CreateClaims(request.BaseClaims{
		Username: "admin",
	})

	token, err := j.CreateToken(claims)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
}
