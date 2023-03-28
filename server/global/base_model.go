package global

import (
	"time"

	"gorm.io/gorm"
)

//重写gorm.Model

type BaseModel struct {
	ID        int64          `json:"id" gorm:"primaryKey;column:id;AUTO_INCREMENT"`
	CreateAt  time.Time      `json:"createAt" gorm:"autoCreateTime;column:create_time;"`       // 使用时间戳秒数填充创建时间
	UpdateAt  time.Time      `json:"UpdateAt" gorm:"autoUpdateTime:milli;column:update_time;"` // 使用时间戳毫秒数填充更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
