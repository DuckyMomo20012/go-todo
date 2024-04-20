package configs

type ServerConfig struct {
	Host              string `mapstructure:"HOST"`
	Port              string `mapstructure:"PORT"`
	TaskServerAddress string `mapstructure:"TASK_SERVER_ADDRESS"`
}
