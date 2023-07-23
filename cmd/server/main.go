package main

import (
	"GSS/internal/server"
	"GSS/internal/server/config"
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"log"
	"net"
)

// TODO сделать методы для вызова их через grpc
type GrpcServer struct {
	pb.UnimplementedGreeterServer
}

func (s *GrpcServer) StateFun(ctx context.Context, in *pb.StateRequest) (*pb.StateReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
	//TODO понять как работает return в grpc
}

func main() {
	cfg, err := config.LoadConfig("./config/server/")
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(cfg)

	server.Run()
	// Код который написн ниже считается говно-кодом)))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &GrpcServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
