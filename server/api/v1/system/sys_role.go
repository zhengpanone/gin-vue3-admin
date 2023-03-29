package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/gin-vue3-admin/model/common/response"
	"github.com/zhengpanone/gin-vue3-admin/model/request"
)

type RoleApi struct{}

// CreateRole
// @Tags RoleApi
// @Summary 创建角色
// @Description 创建角色
// @ID /v1/api/role/register
// @Accept  json
// @Produce  application/json
// @Param data body request.RoleParam true "body" #[roleName,密码,手机号码] body [string,string,string] [required,required,required] "[system.SysUser]"
// @Router /v1/api/role/addRole [post]
func (r *RoleApi) CreateRole(c *gin.Context) {
	var roleParam request.RoleParam
	_ = c.ShouldBindJSON(&roleParam)
	role, err := roleService.CreateRole(roleParam)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
		return
	}
	response.OkWithData(c, role)
}
