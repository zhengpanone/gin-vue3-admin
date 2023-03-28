package config

// JWT JSON WEB TOKEN 配置
type JWT struct {
	SigningKey string `yaml:"signingKey" mapstructure:"signingKey"` // jwt签名
	Issuer     string `yaml:"issuer" mapstructure:"issuer"`         // 签发者
	ExpireTime int64  `yaml:"expire" mapstructure:"expire"`         // 过期时间
	BufferTime int64  `yaml:"bufferTime" mapstructure:"bufferTime"` // 缓冲时间
}
