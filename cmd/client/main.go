package main

import (
	"GSS/internal/client"
	"GSS/internal/client/config"
	"log"
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

	//grpc request state fun
	client.StateFunCli(cfg)

}
