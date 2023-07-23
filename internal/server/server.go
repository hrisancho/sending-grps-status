package server

import (
	"GSS/internal/metrics"
	"GSS/internal/server/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

type Server struct {
	App          *fiber.App
	Config       config.Config
	UsersMetrics map[uuid.UUID][]metrics.MetricStorage
}

func NewServer(config config.Config) (server *Server) {
	server = &Server{
		App:          fiber.New(),
		Config:       config,
		UsersMetrics: make(map[uuid.UUID][]metrics.MetricStorage),
	}
	server.App.Use(logger.New(logger.Config{
		Format: config.Logger_fmt,
	}))

	server.SetupRoutes()

	return
}

func (server *Server) Run() {
	err := server.App.Listen(server.Config.ServerAddr)
	if err != nil {
		log.Fatal(err)
	}
}
