package utils

import (
	"github.com/mojocn/base64Captcha"
)

type CaptchaResult struct {
	Id          string `json:"id"`
	Base64Blog  string `json:"base_64_blog"`
	VerifyValue string `json:"code"`
}

// 默认存储10240个验证码,每个验证码10分钟过期
var store = base64Captcha.DefaultMemStore

// GenerateCaptcha 生产验证码
func (cap *CaptchaResult) GenerateCaptcha() (*CaptchaResult, error) {
	driver := base64Captcha.NewDriverDigit(70, 130, 4, 0.8, 100)
	/*driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor:         &color.RGBA{R: 3, G: 102, B: 214, A: 125},
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	driver = driverString.ConvertFonts()*/

	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		return &CaptchaResult{}, err
	}
	captchaResult := CaptchaResult{Id: id, Base64Blog: b64s}
	return &captchaResult, nil

}

// VerifyCaptcha 验证验证码
func (cap *CaptchaResult) VerifyCaptcha(id string, value string) bool {
	if store.Verify(id, value, true) {
		// 验证成功
		return true
	} else {
		// 验证失败
		return false
	}

}
