package router

type RouterGroup struct {
	BaseRouter
	UserRouter
	JwtRouter
	RoleRouter
	MenuRouter
}

var RouterGroupApp = new(RouterGroup)
