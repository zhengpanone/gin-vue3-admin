package system

import (
	"github.com/zhengpanone/gin-vue3-admin/global"
	"github.com/zhengpanone/gin-vue3-admin/model/entity/system"
)

type OperationRecordService struct{}

func (operationRecordService *OperationRecordService) CreateSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.GVA_DB.Create(&sysOperationRecord).Error
	return err
}
