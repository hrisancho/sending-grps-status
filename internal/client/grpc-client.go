package client

import (
	"GSS/internal/client/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func Connect2Grpc(config config.Config) (cli *grpc.ClientConn) {

	//Подключение к серверу
	conn, err := grpc.Dial(config.GrpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	return conn
}
