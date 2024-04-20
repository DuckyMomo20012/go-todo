package server

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGRPCServer(registerServer func(server *grpc.Server)) {
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
		logrus.Fatal(err)
	}

	logrus.WithField("grpcEndpoint", grpcEndpoint).Info("Starting: gRPC Listener")
	logrus.Fatal(grpcServer.Serve(listen))
}
