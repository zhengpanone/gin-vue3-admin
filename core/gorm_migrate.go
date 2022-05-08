package core

import (
	"fmt"
	"gin-api-learn/global"
	"gin-api-learn/model/entity/system"
	"gorm.io/gorm"
)

func setTableOption(tableComment string) *gorm.DB {
	value := fmt.Sprintf("ENGINE=InnoDB COMMENT='%s'", tableComment)
	return global.GlobalMysqlClient.Set("gorm:table_options", value)

}

func sysTable() {
	// 用户账户表
	_ = setTableOption("用户表").AutoMigrate(&system.SysUser{})
	_ = setTableOption("用户信息表").AutoMigrate(&system.SysUserInfo{})
	_ = setTableOption("角色表").AutoMigrate(&system.SysRole{})
	_ = setTableOption("用户角色表").AutoMigrate(&system.SysUserRole{})
	_ = setTableOption("菜单表").AutoMigrate(&system.SysMenu{})
	_ = setTableOption("角色菜单表").AutoMigrate(&system.SysRoleMenu{})

}
func AutoMigrate() {
	// 创建用户相关表
	sysTable()
}
