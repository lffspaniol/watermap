package grpc

import (
	"fmt"
	"log"
	"net"

	"watermap/cmd/server/providers"
	pb "watermap/gen/proto"
	"watermap/infrastructure/db"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type rpc struct {
	repository db.Repository
}

func (g *rpc) Serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", viper.GetInt("grpc_port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	log.Println("Starting server Listen in port:", viper.GetInt("grpc_port"))
	pb.RegisterGreeterServer(grpcServer, &Greeter{
		repository: g.repository,
	})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start serve: %v", err)
	}
}

func Provider() providers.Server {
	return &rpc{}
}
