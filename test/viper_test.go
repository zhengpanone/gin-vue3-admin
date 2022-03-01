package test

import (
	"gin-api-learn/global"
	"testing"

	"github.com/spf13/viper"
)

func TestViper1(t *testing.T) {
	v := viper.New()
	v.SetConfigFile(global.CONFIGFILE)
	if err := v.ReadInConfig(); err != nil {
		t.Errorf(err.Error())
	}
	v.RegisterAlias("author", "username")
	v.Set("name", "测试")
	err := v.Unmarshal(&global.GlobalConfig)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Logf("config:%#v\n", global.GlobalConfig)
}
