package service

type ServiceGroup struct {
	UserService
	JwtService
	SysRoleService
	MenuService
	OperationRecordService
}

var ServiceGroupApp = new(ServiceGroup)
