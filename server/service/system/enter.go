package system

type ServiceGroup struct {
	UserService
	JwtService
	SysRoleService
	MenuService
	OperationRecordService
}
