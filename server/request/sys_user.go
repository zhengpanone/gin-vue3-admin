package request

import (
	"github.com/zhengpanone/gin-vue3-admin/model"
)

// 参数校验 github.com/go-playground/validator

// 注册参数

type RegisterParam struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 明文加密后的密码
	RoleID   string `json:"roleID" gorm:"default:111"`
	Phone    string `json:"phone" binding:"required"`                            // 手机号码
	NickName string `json:"nickName" binding:"required" validate:"min=3,max=32"` // 昵称
	Birthday string `json:"birthday"`                                            // 生日
	Address  string `json:"address"`                                             // 地址
	Email    string `json:"email"`                                               // 邮箱
}

type ChangePasswordParam struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

type LoginParam struct {
	Username  string `json:"username" binding:"required,gt=2"` // 用户名
	Password  string `json:"password" binding:"required,gt=6"` // 密码
	Captcha   string `json:"captcha"`                          // 验证码
	CaptchaID string `json:"captchaID"`                        // 验证码ID
}

type ChangeUserInfo struct {
	ID           uint                 `gorm:"primarykey"`                                                                           // 主键ID
	NickName     string               `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone        string               `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	AuthorityIds []uint               `json:"authorityIds" gorm:"-"`                                                                // 角色ID
	Email        string               `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	HeaderImg    string               `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	SideMode     string               `json:"sideMode"  gorm:"comment:用户侧边主题"`                                                      // 用户侧边主题
	Enable       int                  `json:"enable" gorm:"comment:冻结用户"`                                                           //冻结用户
	Authorities  []model.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}
