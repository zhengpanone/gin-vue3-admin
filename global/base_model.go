package global

import (
	"time"

	"gorm.io/gorm"
)

//重写gorm.Model

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreateAt  time.Time      `json:"createAt" gorm:"autoCreateTime"`       // 使用时间戳秒数填充创建时间
	UpdateAt  time.Time      `json:"UpdateAt" gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
