package config

// JWT JSON WEB TOKEN 配置
type JWT struct {
	SigningKey string `yaml:"signingKey"` // jwt签名
	Issuer     string `yaml:"issuer"`     // 签发者
	ExpireTime int64  `yaml:"expire"`     // 过期时间
	BufferTime int64  `yaml:"bufferTime"` // 缓冲时间
}
