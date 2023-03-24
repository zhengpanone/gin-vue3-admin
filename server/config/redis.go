package config

import "time"

type redis struct {
	Addr        string        `yaml:"addr"`
	Password    string        `yaml:"password"`
	DefaultDB   int           `yaml:"defaultDB"`
	DialTimeout time.Duration `yaml:"dialinteout"`
}
