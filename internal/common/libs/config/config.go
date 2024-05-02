package config

import (
	"github.com/DuckyMomo20012/go-todo/internal/common/libs/logger"
	"github.com/spf13/viper"
)

func LoadConfig[T any](config *T, path string) error {
	log := logger.Get()

	viper.SetConfigName(".env") // allow directly reading from .env file
	viper.SetConfigType("env")
	viper.AddConfigPath(".")  // optionally look for config in the working directory
	viper.AddConfigPath(path) // call multiple times to add many search paths
	viper.AddConfigPath("/")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Err(err).Msg("failed to read config file")

		return err
	}

	if err := viper.Unmarshal(config); err != nil {
		log.Error().Err(err).Msg("failed to unmarshal config")

		return err
	}

	return nil
}
