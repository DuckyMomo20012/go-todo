package tasks

import (
	"fmt"

	tasksv1 "github.com/DuckyMomo20012/go-todo/internal/common/genproto/tasks/v1"
	cfg "github.com/DuckyMomo20012/go-todo/internal/common/libs/config"
	"github.com/DuckyMomo20012/go-todo/internal/common/server"
	"github.com/DuckyMomo20012/go-todo/internal/tasks/adapters"
	"github.com/DuckyMomo20012/go-todo/internal/tasks/app"
	"github.com/DuckyMomo20012/go-todo/internal/tasks/configs"
	"github.com/DuckyMomo20012/go-todo/internal/tasks/ports"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"xorm.io/xorm"
)

func NewTaskCmd() *cobra.Command {
	tasksCmd := &cobra.Command{
		Use:   "tasks",
		Short: "Tasks server commands.",
		Long:  "Tasks server commands.",
	}

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the tasks server.",
		Long:  "Start the tasks server. This will start the gRPC server for the tasks service.",
		Run: func(cmd *cobra.Command, args []string) {
			startTaskServer()
		},
	}

	tasksCmd.AddCommand(startCmd)

	return tasksCmd
}

func startTaskServer() {
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")

	var config configs.ServerConfig

	if err := cfg.LoadConfig(&config, "./internal/tasks/configs"); err != nil {
		log.Error(fmt.Sprintf("Error loading config, %s", err))
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
	if err := engine.Sync(new(app.Task)); err != nil {
		log.Error(fmt.Sprintf("Error syncing database, %s", err))

		return
	}

	taskRepository := adapters.NewPgTaskRepository(engine)

	taskServer := ports.NewGrpcServer(taskRepository)

	server.RunGRPCServer(func(server *grpc.Server) {
		tasksv1.RegisterTaskServiceServer(server, taskServer)
	})
}
