package server

import (
	"fmt"
	"net"

	"github.com/DuckyMomo20012/go-todo/internal/common/libs/logger"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGRPCServer(registerServer func(server *grpc.Server)) {
	log := logger.Get()

	port := viper.Get("PORT")

	grpcEndpoint := fmt.Sprintf(":%s", port)

	grpcServer := grpc.NewServer()
	registerServer(grpcServer)

	// NOTE: This may be a security risk, as it allows anyone to see the
	// services available on the server. It is recommended to remove this in a
	// production environment.
	// Currently, it is enabled for Postman testing purposes.
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	listen, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Panic().Err(err).Msgf("failed to listen on %s", grpcEndpoint)
	}

	log.Info().Str("grpcEndpoint", grpcEndpoint).Msg("Starting: gRPC Listener")

	if err := grpcServer.Serve(listen); err != nil {
		log.Panic().Err(err).Msg("failed to serve gRPC server")
	}
}
