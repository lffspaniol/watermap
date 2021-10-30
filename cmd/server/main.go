package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/spf13/viper"

	pb "watermap/gen/proto"
	grpcSv "watermap/infrastructure/server/grpc"
)

func main() {
	err := viper.ReadInConfig()
	viper.SetConfigFile("config.yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config file not found, using OS env vars")
	}

	viper.AutomaticEnv()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", viper.GetInt("grpc_port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	log.Println("Starting server Listen in port:", viper.GetInt("grpc_port"))
	pb.RegisterGreeterServer(grpcServer, &grpcSv.Greeter{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start serve: %v", err)
	}
}
