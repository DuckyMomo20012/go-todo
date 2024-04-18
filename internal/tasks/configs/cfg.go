package configs

type ServerConfig struct {
	Host             string `mapstructure:"HOST"`
	Port             string `mapstructure:"PORT"`
	DBHost           string `mapstructure:"DB_HOST"`
	DBPort           string `mapstructure:"DB_PORT"`
	DBUser           string `mapstructure:"DB_USER"`
	DBPass           string `mapstructure:"DB_PASSWORD"`
	DBName           string `mapstructure:"DB_NAME"`
	CorsAllowOrigins string `mapstructure:"CORS_ALLOW_ORIGINS"`
	ServerProtocol   string `mapstructure:"SERVER_PROTOCOL"`
}
