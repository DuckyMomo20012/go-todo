package configs

type ServerConfig struct {
	Host              string `mapstructure:"HOST"`
	Port              string `mapstructure:"PORT"`
	TaskServerAddress string `mapstructure:"TASK_SERVER_ADDRESS"`
	AppEnv            string `mapstructure:"APP_ENV"`
	LogLevel          string `mapstructure:"LOG_LEVEL"`
	LogSampleRate     string `mapstructure:"LOG_SAMPLE_RATE"`
}
