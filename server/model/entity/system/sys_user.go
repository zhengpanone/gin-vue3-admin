package system

import (
	"github.com/gofrs/uuid/v5"
	"github.com/zhengpanone/gin-vue3-admin/global"
)

// SysUser 用户表
type SysUser struct {
	global.BaseModel
	UUID        uuid.UUID      `json:"uuid" gorm:"index;comment:用户UUID"`                                                     // 用户UUID
	Username    string         `json:"userName" gorm:"index;comment:用户登录名"`                                                  // 用户登录名
	Password    string         `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName    string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	SideMode    string         `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                          // 用户侧边主题
	HeaderImg   string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor   string         `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                           // 基础颜色
	ActiveColor string         `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                      // 活跃颜色
	AuthorityId uint           `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                        // 用户角色ID
	Authority   SysAuthority   `json:"authority" gorm:"-"`                                                                   // `gorm:"foreignKey:关联表的结构体字段;references:当前表的结构体字段;`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone       string         `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email       string         `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enable      int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结

	Status   string      `json:"status" gorm:"size:4;default:1;comment:状态 1:正常 2:白名单 3:黑名单"`
	UserInfo SysUserInfo `json:"userInfo" gorm:"-"`
	Token    string      `json:"token" gorm:"-"`
	Sex      string      `validate:"oneof=female male" gorm:"column:sex;comment:性别"`
}

// SysUserInfo 用户信息表
type SysUserInfo struct {
	UserID   string `json:"userID"`
	Birthday string `json:"birthday" gorm:"type:varchar(10);comment:生日"`
	Address  string `json:"address" gorm:"type:text;comment:地址"`
	global.BaseModel
}

func (SysUser) TableName() string {
	return "sys_user"
}
