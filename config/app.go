package config

// 应用信息
type app struct {
	Desc       string `yaml:"desc"`
	Addr       string `yaml:"addr"`
	ConfigFile string `yaml:"configFile"`
	Version    string `yaml:"Version"`
	Env        string `yaml:"env"`
}

// redis
type redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// serverconfig配置信息
type ServerConfig struct {
	App   app   `yaml:"app"`
	Mysql mysql `yaml:"mysql"`
	Redis redis `yaml:"redis"`
	Log   log   `yaml:"log"` //嵌入日志配置
	Jwt   Jwt   `yaml:"jwt"`
}
