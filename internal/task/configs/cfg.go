package configs

import (
	"github.com/DuckyMomo20012/go-todo/internal/common/libs/config"
	"github.com/spf13/viper"
)

type TaskServerConfig struct {
	config.BaseConfig
	DBUrl string `mapstructure:"DB_URL"`
}

func SetDefaultTaskConfig() {
	config.SetDefaultBaseConfig()

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
}
