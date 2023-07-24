package server

import (
	"GSS/internal/server/config"
	pb "GSS/proto/grpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type GrpcServer struct {
	pb.UnimplementedStatusServer
}

func (s *GrpcServer) StateFun(ctx context.Context, in *pb.StateRequest) (*pb.StateReply, error) {
	log.Printf("Received: ", in)
	return &pb.StateReply{Hi: "Hello "}, nil
}

func RunGrpcServer(config config.Config) {

	_, port, err := net.SplitHostPort(config.GrpcServerAddr)
	if err != nil {
		panic(err)
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", portInt))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpc_server := grpc.NewServer()

	pb.RegisterStatusServer(grpc_server, &GrpcServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpc_server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return
}
