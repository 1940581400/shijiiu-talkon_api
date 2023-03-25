package forms

// EmailLogin 邮箱登录表单
type EmailLogin struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

// MobileLogin 手机登录表单
type MobileLogin struct {
	Mobile   string `form:"email" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
