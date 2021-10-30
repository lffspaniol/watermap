package grpc

import (
	"context"

	pb "watermap/gen/proto"
)

type Greeter struct {
	pb.UnimplementedGreeterServer
}

func (s Greeter) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "abc",
	}, nil
}
