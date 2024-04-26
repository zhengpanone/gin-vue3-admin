package initialize

import (
	"fmt"

	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"

	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

/**
* 初始化mysql主文件
 */

// InitGorm 初始化mysql客户端
func InitGorm() {
	mysqlConfig := global.GVA_CONFIG.Mysql
	//user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database, mysqlConfig.Charset, mysqlConfig.ParseTime, mysqlConfig.TimeZone)

	// 设置gorm配置
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: mysqlConfig.Gorm.SkipDefaultTx, // 是否跳过默认事务
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   mysqlConfig.Gorm.TablePrefix,
			SingularTable: mysqlConfig.Gorm.SingularTable,
		},
		// 执行任何SQL时都会创建一个prepared statement并将其缓存，以提高后续的效率
		PrepareStmt: mysqlConfig.Gorm.PreparedStmt,
		//在AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true
		DisableForeignKeyConstraintWhenMigrating: mysqlConfig.Gorm.CloseForeignKey,
	}
	// 是否覆盖默认是sql配置
	if mysqlConfig.Gorm.CoverLogger {
		setNewLogger(gormConfig)
	}
	client, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         mysqlConfig.DefaultStringSize,
		DisableDatetimePrecision:  mysqlConfig.DisableDatetimePrecision,
		SkipInitializeWithVersion: mysqlConfig.SkipInitializeWithVersion,
	}), gormConfig)
	if err != nil {
		panic(fmt.Sprintf("创建mysql客户端失败：%s", err))

	}
	global.GVA_DB = client
	if mysqlConfig.AutoMigrate {
		RegisterTables()
	}
}

func setTableOption(tableComment string) *gorm.DB {
	value := fmt.Sprintf("ENGINE=InnoDB COMMENT='%s'", tableComment)
	return global.GVA_DB.Set("gorm:table_options", value)

}

func RegisterTables() {
	// 创建用户相关表
	sysTable()
}

// 设置新的Logger
func setNewLogger(gConfig *gorm.Config) {
	logPath := global.GVA_CONFIG.Log.Path
	file, _ := os.OpenFile(logPath+"/sql.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	logLevelMap := map[string]logger.LogLevel{
		"error": logger.Error,
		"info":  logger.Info,
		"warn":  logger.Warn,
	}
	var logLevel logger.LogLevel
	var ok bool

	if logLevel, ok = logLevelMap[global.GVA_CONFIG.Mysql.LogLevel]; !ok {
		logLevel = logger.Error
	}
	newLogger := logger.New(log.New(file, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             global.GVA_CONFIG.Mysql.SlowSql,                   //慢SQL时间
		LogLevel:                  logLevel,                                          // 记录日志级别
		IgnoreRecordNotFoundError: global.GVA_CONFIG.Mysql.IgnoreRecordNotFoundError, // 是否忽略ErrRecordNotFound(未查到记录错误)
		Colorful:                  false,                                             // 开关颜色
	})
	gConfig.Logger = newLogger

}

func sysTable() {

	_ = setTableOption("用户表").AutoMigrate(&system.SysUser{})
	_ = setTableOption("用户信息表").AutoMigrate(&system.SysUserInfo{})
	_ = setTableOption("角色表").AutoMigrate(&system.SysAuthority{})
	_ = setTableOption("菜单表").AutoMigrate(&system.SysBaseMenu{})
	_ = setTableOption("菜单参数表").AutoMigrate(&system.SysBaseMenuParameter{})

	_ = setTableOption("角色菜单表").AutoMigrate(&system.SysAuthorityMenu{})
	_ = setTableOption("用户角色表").AutoMigrate(&system.SysUserAuthority{})

	_ = setTableOption("菜单按钮表").AutoMigrate(&system.SysBaseMenuBtn{})

	// 用户账户表
	//_ = setTableOption("用户角色表").AutoMigrate(&system.SysUserRole{})
	//_ = setTableOption("角色权限表").AutoMigrate(&system.SysRolePermission{})

	//_ = setTableOption("jwt黑名单表").AutoMigrate(&system.JwtBlacklist{})

}
