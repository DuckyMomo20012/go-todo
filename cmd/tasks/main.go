package main

import (
	"fmt"

	"github.com/DuckyMomo20012/go-todo/internal/common/server"
	"github.com/DuckyMomo20012/go-todo/internal/tasks/adapters"
	"github.com/DuckyMomo20012/go-todo/internal/tasks/app"
	"github.com/DuckyMomo20012/go-todo/internal/tasks/configs"
	"github.com/DuckyMomo20012/go-todo/internal/tasks/ports"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"xorm.io/xorm"
)

func main() {
	viper.SetConfigName("cfg")
	viper.SetConfigType("env")
	viper.AddConfigPath("./internal/tasks/configs")

	viper.AutomaticEnv()

	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")

	var config *configs.ServerConfig

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	engine, err := xorm.NewEngine("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.DbHost,
			config.DbPort,
			config.DbUser,
			config.DbPass,
			config.DbName,
		))

	if err != nil {
		fmt.Printf("Error creating engine, %s", err)
		return
	}

	// NOTE: Sync database models
	err = engine.Sync(new(app.Task))

	if err != nil {
		fmt.Printf("Error syncing database, %s", err)
		return
	}

	taskRepository := adapters.NewPgTaskRepository(engine)

	taskService := app.NewTaskService(taskRepository)

	httpServer := ports.NewHttpServer(taskService)

	server.RunHttpServer(func(app *fiber.App) {
		ports.RegisterHandlers(app, httpServer)
	})
}
