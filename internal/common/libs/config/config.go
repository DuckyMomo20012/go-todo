package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadConfig[T any](config *T, path string) error {
	viper.SetConfigName(".env") // allow directly reading from .env file
	viper.SetConfigType("env")
	viper.AddConfigPath(".")  // optionally look for config in the working directory
	viper.AddConfigPath(path) // call multiple times to add many search paths
	viper.AddConfigPath("/")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Error(fmt.Sprintf("Error reading config file, %s", err))
	}

	if err := viper.Unmarshal(config); err != nil {
		log.Error(fmt.Sprintf("Error unmarshalling config, %s", err))
	}

	return nil
}
