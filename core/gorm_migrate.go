package core

import (
	"fmt"
	"gin-api-learn/global"
	"gin-api-learn/model/entity"

	"gorm.io/gorm"
)

func setTableOption(tableComment string) *gorm.DB {
	value := fmt.Sprintf("ENGINE=InnoDB COMMENT='%s'", tableComment)
	return global.GlobalMysqlClient.Set("gorm:table_options", value)

}

//用户相关
func userTable() {
	// 用户账户表
	_ = setTableOption("用户表").AutoMigrate(&entity.User{})

	_ = setTableOption("用户信息表").AutoMigrate(&entity.UserInfo{})

}
func AutoMigrate() {
	// 创建用户相关表
	userTable()
}
