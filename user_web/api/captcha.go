package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"net/http"
	"talkon_api/user_web/global"
	"talkon_api/user_web/global/response"
)

func GetPicCaptcha(ctx *gin.Context) {
	digit := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(digit, global.CaptchaStore)
	id, b64s, err := captcha.Generate()
	if err != nil {
		zap.L().Error("[GetPicCaptcha] 生成验证码失败", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "验证码生成失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, response.CaptchaInfo{
		CaptchaId: id,
		B64s:      b64s,
	})
}
