package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunGatewayServer(registerServer func(ctx context.Context, mux *runtime.ServeMux, opts []grpc.DialOption) error) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := registerServer(ctx, mux, opts)
	if err != nil {
		return err
	}

	// Create an HTTP server with desired timeouts
	const timeoutSeconds = 10

	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", viper.Get("PORT")),
		Handler:      mux,
		ReadTimeout:  time.Second * timeoutSeconds, // Set the read timeout to 10 seconds
		WriteTimeout: time.Second * timeoutSeconds, // Set the write timeout to 10 seconds
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.WithField("port", viper.Get("PORT")).Info("Starting: HTTP Listener")
	return server.ListenAndServe()
}
