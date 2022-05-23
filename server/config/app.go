package config

// 应用信息
type app struct {
	Desc       string `yaml:"desc"`
	Addr       string `yaml:"addr"`
	ConfigFile string `yaml:"configFile"`
	Version    string `yaml:"Version"`
	Env        string `yaml:"env"`
}

// ServerConfig serverConfig 配置信息
type ServerConfig struct {
	App     app     `yaml:"app" mapstructure:"app"`
	Mysql   mysql   `yaml:"mysql" mapstructure:"mysql"`
	Redis   redis   `yaml:"redis" mapstructure:"redis"`
	Log     log     `yaml:"log" mapstructure:"log"` //嵌入日志配置
	JWT     JWT     `yaml:"jwt" mapstructure:"jwt"`
	Captcha Captcha `yaml:"captcha" mapstructure:"captcha"`
}
