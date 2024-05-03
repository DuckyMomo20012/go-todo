package gateway

import (
	"context"
	"flag"

	taskv1 "github.com/DuckyMomo20012/go-todo/internal/common/genproto/task/v1"
	cfg "github.com/DuckyMomo20012/go-todo/internal/common/libs/config"
	"github.com/DuckyMomo20012/go-todo/internal/common/libs/logger"
	"github.com/DuckyMomo20012/go-todo/internal/common/server"
	"github.com/DuckyMomo20012/go-todo/internal/gateway/configs"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.uber.org/multierr"
	"google.golang.org/grpc"
)

func NewGatewayCmd() *cobra.Command {
	gatewayCmd := &cobra.Command{
		Use:   "gateway",
		Short: "Gateway server commands.",
		Long:  "Gateway server commands.",
	}

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the gateway server.",
		Long:  "Start the gateway server. This will start the HTTP gateway proxy server.",
		Run: func(cmd *cobra.Command, args []string) {
			startGatewayServer()
		},
	}

	gatewayCmd.AddCommand(startCmd)

	return gatewayCmd
}

func startGatewayServer() {
	var config configs.GatewayServerConfig

	configs.SetDefaultGatewayServerConfig()

	if err := cfg.LoadConfig(&config, "./internal/gateway/configs"); err != nil {
		log.Error().Err(err).Msg("failed to load config")
	}

	// NOTE: Load config before setting up logger
	logger.Get()
	logger.SetService("gateway")

	taskServerEndpoint := flag.String("task-server-endpoint", config.TaskServerAddress, "task gRPC server endpoint")

	server.RunGatewayServer(func(ctx context.Context, mux *runtime.ServeMux, opts []grpc.DialOption) error {
		err := multierr.Combine(
			taskv1.RegisterTaskServiceHandlerFromEndpoint(ctx, mux, *taskServerEndpoint, opts),
		)
		if err != nil {
			return err
		}

		return nil
	})
}
