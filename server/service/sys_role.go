package service

import (
	"errors"
	"github.com/zhengpanone/gin-vue3-admin/dao"
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model"
	"github.com/zhengpanone/gin-vue3-admin/request"
	"gorm.io/gorm"
)

type SysRoleService struct{}

func (roleService *SysRoleService) CreateRole(r request.RoleParam) (role model.SysRole, err error) {
	roleDao := dao.SysRoleDao{}
	// 查找父级id是否存在
	var parentRole model.SysRole
	if errors.Is(global.GVA_DB.Where("role_code=?", r.ParentCode).First(&parentRole).Error, gorm.ErrRecordNotFound) {
		r.ParentCode = "0"
	}
	role = model.SysRole{
		RoleName: r.RoleName,
		RoleDesc: r.RoleDesc,
		RoleCode: r.RoleCode,
	}
	role, err = roleDao.AddRole(role)

	return role, err
}

func (roleService *SysRoleService) PageRole(info request.PageRequest) (list []model.SysRole, total int64, err error) {
	limit := info.PageSize
	offset := limit * (info.Page - 1)
	db := global.GVA_DB.Model(&model.SysRole{})
	if err = db.Where("parent_id=?", "0").Count(&total).Error; total == 0 || err != nil {
		return
	}
	var role []model.SysRole
	err = db.Limit(limit).Offset(offset).Where("parent_id=?", "0").Find(&role).Error
	limit = len(role)
	for k := range role {
		roleService.findChildrenRole(&role[k])
	}
	return role, total, err
}

func (roleService *SysRoleService) findChildrenRole(role *model.SysRole) {
	_ = global.GVA_DB.Where("parent_code=?", role.RoleCode).First(&role.Children).Error
	if len(role.Children) > 0 {
		for k := range role.Children {
			roleService.findChildrenRole(&role.Children[k])
		}
	}
}

func (roleService *SysRoleService) DeleteRole(id string) error {
	var data []model.SysRole
	global.GVA_DB.Model(&model.SysRole{}).Where("id=?", id).Find(&data)
	if len(data) == 0 {
		return errors.New("id没有找到，删除失败")
	}
	global.GVA_DB.Delete(&data)
	return nil
}

func (roleService *SysRoleService) UpdateRole() error {
	return nil
}
