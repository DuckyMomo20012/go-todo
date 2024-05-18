package task

import (
	taskv1 "github.com/DuckyMomo20012/go-todo/internal/common/genproto/task/v1"
	cfg "github.com/DuckyMomo20012/go-todo/internal/common/libs/config"
	"github.com/DuckyMomo20012/go-todo/internal/common/libs/db"
	"github.com/DuckyMomo20012/go-todo/internal/common/libs/logger"
	"github.com/DuckyMomo20012/go-todo/internal/common/server"
	"github.com/DuckyMomo20012/go-todo/internal/task/adapters"
	"github.com/DuckyMomo20012/go-todo/internal/task/configs"
	"github.com/DuckyMomo20012/go-todo/internal/task/ports"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
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
	var config configs.TaskServerConfig

	configs.SetDefaultTaskConfig()

	if err := cfg.LoadConfig(&config, "./internal/task/configs"); err != nil {
		log.Error().Err(err).Msg("failed to load config")
	}

	// NOTE: Load config before setting up logger
	logger.Get()
	logger.WithService("task")

	dbpool := db.NewDb(config.DBUrl)

	defer dbpool.Close()

	taskRepository := adapters.NewPgTaskRepository(dbpool)

	taskServer := ports.NewGrpcServer(taskRepository)

	server.RunGRPCServer(func(server *grpc.Server) {
		taskv1.RegisterTaskServiceServer(server, taskServer)
	})
}
