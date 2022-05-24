package test

import (
	"gin-api-learn/server/global"
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
	err := v.Unmarshal(&global.GVA_CONFIG)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Logf("config:%#v\n", global.GVA_CONFIG)
}
