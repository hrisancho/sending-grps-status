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

	go func() {
		serv := server.NewWebServer(cfg)
		// Start the web server
		serv.Run()
	}()
	//serv := server.NewWebServer(cfg)
	//
	//serv.Run()

	go func() {
		server.RunGrpcServer(cfg)
	}()

	select {}
}
