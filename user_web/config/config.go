package config

type ServerConfig struct {
	Name          string        `mapstructure:"name"`
	UserWebInfo   UserWebConfig `mapstructure:"server"`
	LoggerInfo    LoggerConfig  `mapstructure:"logger"`
	UserSrvConfig UserSrvConfig `mapstructure:"user_srv"`
	JwtInfo       JwtInfo       `mapstructure:"jwt"`
}

type UserWebConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

// LoggerConfig struct 的配置字段含义请参考 zap.Config
type LoggerConfig struct {
	Level            string `mapstructure:"level"`
	Development      string `mapstructure:"development"`
	Encoding         string `mapstructure:"encoding"`
	EncoderConfig    string `mapstructure:"encoderConfig"`
	OutputPaths      string `mapstructure:"outputPaths"`
	ErrorOutputPaths string `mapstructure:"errorOutputPaths"`
}

// JwtInfo 字段意义参考 jwt.RegisteredClaims
type JwtInfo struct {
	Key      string `mapstructure:"key"`
	Expires  int64  `mapstructure:"expires"`
	Issuer   string `mapstructure:"issuer"`
	Subject  string `mapstructure:"subject"`
	Audience string `mapstructure:"audience"`
	Platform string `mapstructure:"platform"`
}
