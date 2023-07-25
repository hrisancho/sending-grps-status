package main

import (
	"GSS/internal/client"
	"GSS/internal/client/config"
	pb "GSS/proto/grpc"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	cfg, err := config.LoadConfig("./config/client")
	if err != nil {
		log.Fatal(err)
	}
	//
	//user, err := client.NewUser(cfg)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = user.StreamingMetrics()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//Подключение к grpc серверу
	conn, err := grpc.Dial(cfg.GrpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewStatusClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//Объявление сообщения
	meseg := client.StateRequest()
	//Отправка изображения
	r, err := c.StateFun(ctx, meseg)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: ", r)
}
