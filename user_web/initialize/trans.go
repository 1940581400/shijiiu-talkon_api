package initialize

import (
	"github.com/gin-gonic/gin/binding"
	enT "github.com/go-playground/locales/en"
	zhT "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"talkon_api/user_web/global"
)

func InitTrans(locale string) {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		ZH := zhT.New()           // 中文翻译器
		EN := enT.New()           // 英文翻译器
		UNI := ut.New(EN, ZH, EN) // 第一个为备用语言环境
		trans, ok := UNI.GetTranslator(locale)
		deftT, _ := UNI.GetTranslator("en")
		if !ok {
			zap.L().Error("[InitTrans] 不存在的翻译器,将使用默认的英文翻译器", zap.String("不存在的翻译器为", locale))
			trans = deftT
		}
		var err error
		switch locale {
		case "zh":
			err = zh.RegisterDefaultTranslations(validate, trans)
		case "en":
			err = en.RegisterDefaultTranslations(validate, trans)
		default:
			err = en.RegisterDefaultTranslations(validate, trans)
		}
		if err != nil {
			zap.L().Error("[InitTrans] 翻译器注册失败", zap.Error(err))
		}
		// 设置全局翻译器
		global.Trans = trans
	}

}
