package global

import (
	"time"

	"gorm.io/gorm"
)

//重写gorm.Model

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreateAt  time.Time      `json:"createAt"`
	UpdateAt  time.Time      `json:"UpdateAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
