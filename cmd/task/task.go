package task

import (
	"context"
	"fmt"

	taskv1 "github.com/DuckyMomo20012/go-todo/internal/common/genproto/task/v1"
	cfg "github.com/DuckyMomo20012/go-todo/internal/common/libs/config"
	"github.com/DuckyMomo20012/go-todo/internal/common/libs/logger"
	"github.com/DuckyMomo20012/go-todo/internal/common/server"
	"github.com/DuckyMomo20012/go-todo/internal/task/adapters"
	"github.com/DuckyMomo20012/go-todo/internal/task/configs"
	"github.com/DuckyMomo20012/go-todo/internal/task/ports"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func NewTaskCmd() *cobra.Command {
	taskCmd := &cobra.Command{
		Use:   "task",
		Short: "Task server commands.",
		Long:  "Task server commands.",
	}

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the task server.",
		Long:  "Start the task server. This will start the gRPC server for the task service.",
		Run: func(cmd *cobra.Command, args []string) {
			startTaskServer()
		},
	}

	taskCmd.AddCommand(startCmd)

	return taskCmd
}

func startTaskServer() {
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "9000")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")

	log := logger.Get()
	logger.SetService("task")

	var config configs.ServerConfig

	if err := cfg.LoadConfig(&config, "./internal/task/configs"); err != nil {
		log.Error().Err(err).Msg("failed to load config")
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Panic().Err(err).Msg("failed to connect to database")
	}
	defer dbpool.Close()

	taskRepository := adapters.NewPgTaskRepository(dbpool)

	taskServer := ports.NewGrpcServer(taskRepository)

	server.RunGRPCServer(func(server *grpc.Server) {
		taskv1.RegisterTaskServiceServer(server, taskServer)
	})
}
