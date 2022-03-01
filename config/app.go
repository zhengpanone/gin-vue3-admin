package config

// 应用信息
type app struct {
	Desc       string `yaml:"desc"`
	Addr       string `yaml:"addr"`
	ConfigFile string `yaml:"configFile"`
	Version    string `yaml:"Version"`
	Env        string `yaml:"env"`
}

// mysql信息
type mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
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
	// Log   log   `yaml:"log"`
}
