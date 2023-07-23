package main

import (
	"GSS/internal/client"
	"GSS/internal/client/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"log"
)

var addr = "localhost:50051"

func main() {
	cfg, err := config.LoadConfig("./config/client")
	if err != nil {
		log.Fatal(err)
	}

	user, err := client.NewUser(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = user.StreamingMetrics()
	if err != nil {
		log.Fatal(err)
	}
	// Код который написн ниже считается говно-кодом)))

	//Подключение к серверу
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	meseg = &pb.StateRequest{
		Name: *name,
	}
}
