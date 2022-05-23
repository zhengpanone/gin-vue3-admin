package config

type Captcha struct {
	KeyLong   int `yaml:"key-long" mapstructure:"key-long"`     // 	验证码长度
	ImgWidth  int `yaml:"img-width" mapstructure:"img-width"`   // 验证码宽度
	ImgHeight int `yaml:"img-height" mapstructure:"img-height"` // 验证码高度
}
