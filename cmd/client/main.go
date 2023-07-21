package main

import (
	"log"

	"GSS/internal/client"
	"GSS/internal/client/config"
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
}
