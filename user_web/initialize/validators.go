package initialize

import (
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"talkon_api/user_web/global"
	customs "talkon_api/user_web/validator"
)

// InitValidators 注册自定义的验证器
func InitValidators() {
	var err error
	defer func(err error) {
		if err != nil {
			zap.L().Panic("自定义验证器注册失败", zap.String("msg", err.Error()))
		}
	}(err)
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = validate.RegisterValidation("mobile", customs.ValidateMobile)
		err = validate.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 手机号码格式错误", true)
		},
			func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("mobile", fe.Field())
				return t
			})
		if err != nil {
			return
		}
	}
}
