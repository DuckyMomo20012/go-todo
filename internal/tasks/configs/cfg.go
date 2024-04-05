package configs

type ServerConfig struct {
	Host             string `mapstructure:"HOST"`
	Port             string `mapstructure:"PORT"`
	DbHost           string `mapstructure:"DB_HOST"`
	DbPort           string `mapstructure:"DB_PORT"`
	DbUser           string `mapstructure:"DB_USER"`
	DbPass           string `mapstructure:"DB_PASSWORD"`
	DbName           string `mapstructure:"DB_NAME"`
	CorsAllowOrigins string `mapstructure:"CORS_ALLOW_ORIGINS"`
}
