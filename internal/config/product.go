package config

import (
	"time"

	"github.com/spf13/viper"
)

type ProductConfig struct {
	DSN               string        `mapstructure:"DSN"`
	HttpServerAddress string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	Sign              string        `mapstructure:"SIGN"`
	Admin             string        `mapstructure:"ADMIN"`
	TokenDuration     time.Duration `mapstructure:"TOKEN_DURATION"`
}

func LoadProductConfig(path string) (config ProductConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("product")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
