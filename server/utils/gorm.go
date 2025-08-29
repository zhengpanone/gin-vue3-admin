package utils

import (
	"github.com/zhengpanone/gin-vue3-admin/request"
	"github.com/zhengpanone/gin-vue3-admin/response"
	"gorm.io/gorm"
)

func Paginate[T any](db *gorm.DB, req request.PageRequest) (response.PageResult[T], error) {

	var result response.PageResult[T]
	var list []T
	var total int64

	// 统计总数
	if err := db.Count(&total).Error; err != nil {
		return result, err
	}
	result.Total = total
	// 分页查询
	offset := (req.Page - 1) * req.PageSize
	if err := db.Limit(req.PageSize).Offset(offset).First(&list).Error; err != nil {
		return result, err
	}
	result.List = list
	return result, nil
}
