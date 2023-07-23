package main

import (
	"GSS/proto/grpc"
	"flag"
	"fmt"
	"log"
	"net"

	pb "google.golang.org/grpc/examples/helloworld/helloworld"

	"GSS/internal/server"
	"GSS/internal/server/config"
)

func main() {
	cfg, err := config.LoadConfig("./config/server/")
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(cfg)

	server.Run()
	// Код который написн ниже считается говно-кодом)))
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
