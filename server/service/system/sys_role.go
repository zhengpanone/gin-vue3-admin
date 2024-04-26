package system

import (
	"errors"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/common/request"
	"github.com/zhengpanone/gin-vue3-admin/model/dao"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
	systemReq "github.com/zhengpanone/gin-vue3-admin/model/system/request"
	"gorm.io/gorm"
)

type SysRoleService struct{}

func (roleService *SysRoleService) CreateRole(r systemReq.RoleParam) (role system.SysRole, err error) {
	roleDao := dao.SysRoleDao{}
	// 查找父级id是否存在
	var parentRole system.SysRole
	if errors.Is(global.GVA_DB.Where("role_id=?", r.ParentId).First(&parentRole).Error, gorm.ErrRecordNotFound) {
		r.ParentId = "0"
	}
	role = system.SysRole{
		RoleName: r.RoleName,
		Remark:   r.Remark,
		ParentId: r.ParentId,
	}
	role, err = roleDao.AddRole(role)

	return role, err
}

func (roleService *SysRoleService) PageRole(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysRole{})
	if err = db.Where("parent_id=?", "0").Count(&total).Error; total == 0 || err != nil {
		return
	}
	var role []system.SysRole
	err = db.Limit(limit).Offset(offset).Where("parent_id=?", "0").Find(&role).Error
	info.PageSize = len(role)
	for k := range role {
		roleService.findChildrenRole(&role[k])
	}
	return role, total, err
}

func (roleService *SysRoleService) findChildrenRole(role *system.SysRole) {
	_ = global.GVA_DB.Where("parent_id=?", role.RoleID).First(&role.Children).Error
	if len(role.Children) > 0 {
		for k := range role.Children {
			roleService.findChildrenRole(&role.Children[k])
		}
	}
}

func (roleService *SysRoleService) DeleteRole(id string) error {
	var data []system.SysRole
	global.GVA_DB.Model(&system.SysRole{}).Where("id=?", id).Find(&data)
	if len(data) == 0 {
		return errors.New("id没有找到，删除失败")
	}
	global.GVA_DB.Delete(&data)
	return nil
}

func (roleService *SysRoleService) UpdateRole() error {
	return nil
}
