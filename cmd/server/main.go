package main

import (
	pb "GSS/proto/grpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcServer struct {
	pb.UnimplementedStatusServer
}

func (s *GrpcServer) StateFun(ctx context.Context, in *pb.StateRequest) (*pb.StateReply, error) {
	log.Printf("Received: ", in)
	return &pb.StateReply{Hi: "Hello "}, nil
}

func main() {
	//cfg, err := config.LoadConfig("./config/server/")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//server := server.NewServer(cfg)
	//
	//server.Run()
	// Код который написн ниже считается говно-кодом)))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStatusServer(s, &GrpcServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
