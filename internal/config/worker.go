package config

import "github.com/spf13/viper"

type WorkerConfig struct {
	GrpcWorkerServerAddress string `mapstructure:"GRPC_WORKER_SERVER_ADDRESS"`
	RedisAddr               string `mapstructure:"REDIS_ADDR"`
	DSN                     string `mapstructure:"DSN"`
	EmailSenderName         string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderPassword     string `mapstructure:"EMAIL_SENDER_PASSWORD"`
	EmailSenderAddress      string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	WorkerHealthz           string `mapstructure:"WORKER_HEALTHZ"`
}

func LoadWorkerConfig(path string) (config WorkerConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("worker")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
