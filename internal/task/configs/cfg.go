package configs

import (
	"github.com/DuckyMomo20012/go-todo/internal/common/libs/config"
	"github.com/spf13/viper"
)

type TaskServerConfig struct {
	config.BaseConfig
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

func SetDefaultTaskConfig() {
	config.SetDefaultBaseConfig()

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
}
