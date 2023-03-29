package system

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zhengpanone/gin-vue3-admin/model/dao"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
	"github.com/zhengpanone/gin-vue3-admin/model/request"
)

type SysRoleService struct{}

func (roleService *SysRoleService) CreateRole(r request.RoleParam) (role system.SysRole, err error) {
	roleDao := dao.SysRoleDao{}
	role = system.SysRole{
		RoleName: r.RoleName,
		Remark:   r.Remark,
		ParentId: uuid.FromStringOrNil(r.ParentId),
	}
	role, err = roleDao.AddRole(role)

	return role, err
}
