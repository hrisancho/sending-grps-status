package main

import (
	"GSS/internal/server"
	"GSS/internal/server/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("./config/server/")
	if err != nil {
		log.Fatal(err)
	}
	//
	//server := server.NewWebServer(cfg)
	//
	//server.Run()
	server.RunGrpcServer(cfg)

}
