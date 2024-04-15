package main

import (
	"fmt"
	"strings"

	tasksv1 "github.com/DuckyMomo20012/go-todo/internal/common/genproto/tasks/v1"
	"github.com/DuckyMomo20012/go-todo/internal/common/server"
	"github.com/DuckyMomo20012/go-todo/internal/tasks/adapters"
	"github.com/DuckyMomo20012/go-todo/internal/tasks/app"
	"github.com/DuckyMomo20012/go-todo/internal/tasks/configs"
	"github.com/DuckyMomo20012/go-todo/internal/tasks/ports"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
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
	viper.SetDefault("CORS_ALLOW_ORIGINS", "*")
	viper.SetDefault("SERVER_TO_RUN", "http")

	var config *configs.ServerConfig

	if err := viper.ReadInConfig(); err != nil {
		log.Error(fmt.Sprintf("Error reading config file, %s", err))
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Error(fmt.Sprintf("Error unmarshalling config, %s", err))
	}

	engine, err := xorm.NewEngine("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.DBHost,
			config.DBPort,
			config.DBUser,
			config.DBPass,
			config.DBName,
		))
	if err != nil {
		log.Error(fmt.Sprintf("Error connecting to database, %s", err))

		return
	}

	// NOTE: Sync database models
	err = engine.Sync(new(app.Task))
	if err != nil {
		log.Error(fmt.Sprintf("Error syncing database, %s", err))

		return
	}

	taskRepository := adapters.NewPgTaskRepository(engine)

	serverType := strings.ToLower(viper.Get("SERVER_TO_RUN").(string))

	switch serverType {
	case "http":
		{
			taskService := app.NewTaskService(taskRepository)

			httpServer := ports.NewHTTPServer(taskService)

			server.RunHTTPServer(func(app *fiber.App) {
				ports.RegisterHandlers(app, httpServer)
			})
		}

	case "grpc":
		{
			taskServer := ports.NewGrpcServer(taskRepository)

			server.RunGRPCServer(func(server *grpc.Server) {
				tasksv1.RegisterTaskServiceServer(server, taskServer)
			})
		}

	default:
		panic("Unsupported server type")
	}
}
