package service

import (
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model"
)

type OperationRecordService struct{}

func (operationRecordService *OperationRecordService) CreateSysOperationRecord(sysOperationRecord model.SysOperationRecord) (err error) {
	err = global.GVA_DB.Create(&sysOperationRecord).Error
	return err
}
