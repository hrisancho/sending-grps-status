package client

import (
	"GSS/internal/client/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	UUID   uuid.UUID
	Client *fiber.Client
	Config config.Config
}

func NewUser(config config.Config) (user *User, err error) {
	user_uuid, err := uuid.Parse(config.UUID)
	log.Printf("Init user(uuid): %s", user_uuid)
	user = &User{
		UUID:   user_uuid,
		Client: fiber.AcquireClient(),
		Config: config,
	}
	return
}
