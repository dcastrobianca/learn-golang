package server

import (
	"context"

	pb "github.com/dcastrobianca/grpc/grpc/generated_code"
	"google.golang.org/grpc"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func CreateServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterServer{})
	return s
}
