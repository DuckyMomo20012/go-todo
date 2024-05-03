package config

import "github.com/spf13/viper"

type BaseConfig struct {
	Host          string `mapstructure:"HOST"`
	Port          string `mapstructure:"PORT"`
	AppEnv        string `mapstructure:"APP_ENV"`
	LogLevel      string `mapstructure:"LOG_LEVEL"`
	LogSampleRate string `mapstructure:"LOG_SAMPLE_RATE"`
}

func SetDefaultBaseConfig() {
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8080")

	viper.SetDefault("APP_ENV", "development")

	viper.SetDefault("LOG_LEVEL", "0")
	viper.SetDefault("LOG_SAMPLE_RATE", "5")
}
