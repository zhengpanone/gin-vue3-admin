package utils

import (
	"github.com/zhengpanone/gin-api-learn/server/global"
	"os"
)

// DirExist 判断目录是否存在,存在返回true
func DirExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateDir 创建目录
func CreateDir(path string) error {

	dirExist, err := DirExist(path)
	if err != nil {
		return err
	}
	if !dirExist {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			global.GVA_LOG.Sugar().Debugf("创建目录%s失败:%s", path, err)
		}
	}
	return err
}
