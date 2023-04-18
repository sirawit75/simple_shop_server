package config

import (
	"time"

	"github.com/spf13/viper"
)

type UserConfig struct {
	DSN                     string        `mapstructure:"DSN"`
	HttpServerAddress       string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	Sign                    string        `mapstructure:"SIGN"`
	TokenDuration           time.Duration `mapstructure:"TOKEN_DURATION"`
	GrpcLoggerServerAddress string        `mapstructure:"GRPC_LOGGER_SERVER_ADDRESS"`
	GrpcWorkerServerAddress string        `mapstructure:"GRPC_WORKER_SERVER_ADDRESS"`
	RedisAddr               string        `mapstructure:"REDIS_ADDR"`
}

func LoadUserConfig(path string) (config UserConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("user")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
