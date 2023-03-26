package forms

// CheckCaptcha 校验验证码需要传入的字段
type CheckCaptcha struct {
	// CaptchaId 验证码id
	CaptchaId string `form:"captchaId" json:"captchaId" binding:"required"`
	// Answer 输入的码值
	Answer string `form:"answer" json:"Answer" binding:"required,min=4,max=4"`
}
