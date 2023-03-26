package response

// CaptchaInfo 生成的验证码信息
type CaptchaInfo struct {
	// CaptchaId 验证码id
	CaptchaId string `json:"captchaId"`
	// B64s 生成的base64码值
	B64s string `json:"b64s"`
}
