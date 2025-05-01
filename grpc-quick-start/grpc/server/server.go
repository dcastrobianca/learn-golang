package server

import (
	"context"

	"example.com/greetings"
	pb "github.com/dcastrobianca/grpc/grpc/generated_code"
	"google.golang.org/grpc"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	msg, _ := greetings.Hello(req.GetName())
	return &pb.HelloReply{Message: msg}, nil
}

func CreateServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterServer{})
	return s
}
