package global

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"talkon_api/user_web/config"
	"talkon_api/user_web/proto"
)

var (
	// ServerConfig 全局配置
	ServerConfig *config.ServerConfig
	// Trans 全局翻译器
	Trans ut.Translator
	// CaptchaStore 内存验证码仓库
	CaptchaStore = base64Captcha.DefaultMemStore
)

// NewUserSrvClient 创建一个用户服务的grpc客户端
func NewUserSrvClient() proto.UserClient {
	host := ServerConfig.UserSrvConfig.Host
	port := ServerConfig.UserSrvConfig.Port
	inse := insecure.NewCredentials()
	conn := grpc.WithTransportCredentials(inse)
	dial, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), conn)
	if err != nil {
		zap.S().Errorf("[用户服务] 连接失败")
	}
	client := proto.NewUserClient(dial)
	return client
}
