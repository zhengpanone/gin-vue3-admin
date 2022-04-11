package request

// 参数校验 github.com/go-playground/validator

// 注册参数

type RegisterParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	NickName string `json:"nickName" binding:"required" validate:"min=3,max=32"`
	Birthday string `json:"birthday"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}

type ChangePasswordStruct struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password""`   // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

type LoginParam struct {
	Username string `json:"username" binding:"required,gt=2"` // 用户名
	Password string `json:"password" binding:"required,gt=6"` // 密码
}
