package gateway

import (
	"context"
	"flag"
	"fmt"

	tasksv1 "github.com/DuckyMomo20012/go-todo/internal/common/genproto/tasks/v1"
	cfg "github.com/DuckyMomo20012/go-todo/internal/common/libs/config"
	"github.com/DuckyMomo20012/go-todo/internal/common/server"
	"github.com/DuckyMomo20012/go-todo/internal/gateway/configs"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8081")

	var config configs.ServerConfig

	if err := cfg.LoadConfig(&config, "./internal/gateway/configs"); err != nil {
		log.Error(fmt.Sprintf("Error loading config, %s", err))
	}

	taskServerEndpoint := flag.String("task-server-endpoint", config.TaskServerAddress, "task gRPC server endpoint")

	err := server.RunGatewayServer(func(ctx context.Context, mux *runtime.ServeMux, opts []grpc.DialOption) error {
		err := multierr.Combine(
			tasksv1.RegisterTaskServiceHandlerFromEndpoint(ctx, mux, *taskServerEndpoint, opts),
		)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Error(fmt.Sprintf("Error running gateway server, %s", err))
	}
}
