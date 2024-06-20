package server

import (
	"github.com/gofiber/fiber/v2"

	"fiber-api/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "fiber-api",
			AppName:      "fiber-api",
		}),

		db: database.New(),
	}

	return server
}
