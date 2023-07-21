package main

import (
	"context"
	"flag"
	"log"
	"time"

	"GSS/internal/client"
	"GSS/internal/client/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

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
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.Hi(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
}
