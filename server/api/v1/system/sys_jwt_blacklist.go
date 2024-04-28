package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/common/response"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
	"github.com/zhengpanone/gin-vue3-admin/utils"
	"go.uber.org/zap"
)

type JwtApi struct{}

// JsonInBlacklist
//
//	@Tags		Jwt
//	@Summary	jwt加入黑名单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Success	200	{object}	response.Response{msg=string}	"jwt加入黑名单"
//	@Router		/v1/api/jwt/jsonInBlacklist [post]
func (j *JwtApi) JsonInBlacklist(c *gin.Context) {
	token, _ := utils.GetToken(c)
	jwt := system.JwtBlacklist{Jwt: token}
	err := jwtService.JsonInBlacklist(jwt)
	if err != nil {
		global.GVA_LOG.Error("jwt作废", zap.Error(err))
		response.ErrorWithMsg(c, "jwt作废失败")
		return
	}
	utils.ClearToken(c)
	response.OkWithMsg(c, "jwt作废成功")
}
