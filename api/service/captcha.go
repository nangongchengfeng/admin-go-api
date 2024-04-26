package service

import (
	"admin-go-api/common/util"
	"image/color"

	"github.com/mojocn/base64Captcha"
)

/**
 * @Author: 南宫乘风
 * @Description: 验证码
 * @File:  captcha.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-25 17:41
 */

// 使用redis作为store
var store = util.RedisStore{}

// CaptMake 生成验证码
// 返回id, b64s字符串
func CaptMake() (id, b64s string) {
	// 创建一个base64Captcha.Driver
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString
	// 配置验证信息
	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	// 将配置信息转换为driverString
	driverString = captchaConfig
	// 使用driverString创建一个验证码
	driver = driverString.ConvertFonts()
	// 创建 base64Captcha 实例：在实例化 base64Captcha 对象时，传入实现了 Store 接口的 RedisStore 实例。
	captcha := base64Captcha.NewCaptcha(driver, store)

	// 生成一个验证码，并返回id, b64s字符串
	// 生成验证码：调用 Generate 方法时，如果内部逻辑正确无误，Set 方法将被调用，验证码数据及其 ID 将存入 Redis。
	lid, lb64s, _ := captcha.Generate()

	return lid, lb64s
}

// CaptVerify 验证captcha是否正确
func CaptVerify(id string, capt string) bool {
	if store.Verify(id, capt, false) {
		return true
	} else {
		return false
	}
}
