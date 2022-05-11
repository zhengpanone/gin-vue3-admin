package request

// 参数校验 github.com/go-playground/validator

// 注册参数

type RegisterParam struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` //密码
	RoleID   string `json:"roleID" gorm:"default:111"`
	Phone    string `json:"phone" binding:"required"`                            // 手机号码
	NickName string `json:"nickName" binding:"required" validate:"min=3,max=32"` // 昵称
	Birthday string `json:"birthday"`                                            // 生日
	Address  string `json:"address"`                                             // 地址
	Email    string `json:"email"`                                               // 邮箱
}

type ChangePasswordParam struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password""`   // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

type LoginParam struct {
	Username string `json:"username" binding:"required,gt=2"` // 用户名
	Password string `json:"password" binding:"required,gt=6"` // 密码
}
