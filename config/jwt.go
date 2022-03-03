package config

import "time"

// JSON WEB TOKEN 配置
type Jwt struct {
	Secret string        `yaml:"secret"`
	Issuer string        `yaml:"issuer"`
	Expire time.Duration `yaml:"expire"`
}
