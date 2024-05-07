package configs

import (
	"github.com/DuckyMomo20012/go-todo/internal/common/libs/config"
	"github.com/spf13/viper"
)

type GatewayServerConfig struct {
	config.BaseConfig
	TaskServerAddress string `mapstructure:"TASK_SERVER_ADDRESS"`
}

func SetDefaultGatewayServerConfig() {
	config.SetDefaultBaseConfig()

	viper.SetDefault("TASK_SERVER_ADDRESS", "localhost:8081")
}
