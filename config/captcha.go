package config

type Captcha struct {
	KeyLong   int `json:"key-long" yaml:"key-long"`     // 	验证码长度
	ImgWidth  int `json:"img-width" yaml:"img-width"`   // 验证码宽度
	ImgHeight int `json:"img-height" yaml:"img-height"` // 验证码高度
}
