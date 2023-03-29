package system

type RouterGroup struct {
	BaseRouter
	UserRouter
	JwtRouter
	RoleRouter
}
