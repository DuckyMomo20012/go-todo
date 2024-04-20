package server

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func RunGRPCServer(registerServer func(server *grpc.Server)) {
	port := viper.Get("PORT")

	grpcEndpoint := fmt.Sprintf(":%s", port)

	grpcServer := grpc.NewServer()
	registerServer(grpcServer)

	listen, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.WithField("grpcEndpoint", grpcEndpoint).Info("Starting: gRPC Listener")
	logrus.Fatal(grpcServer.Serve(listen))
}
