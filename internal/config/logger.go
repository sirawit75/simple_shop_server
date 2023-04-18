package config

import "github.com/spf13/viper"

type LoggerConfig struct {
	GrpcLoggerServerAddress string `mapstructure:"GRPC_LOGGER_SERVER_ADDRESS"`
	DSN                     string `mapstructure:"DSN"`
}

func LoadLoggerConfig(path string) (config LoggerConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("logger")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
