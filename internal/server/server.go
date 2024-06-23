package server

import (
	"github.com/gofiber/fiber/v2"

	"blueprint-project/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "blueprint-project",
			AppName:      "blueprint-project",
		}),

		db: database.New(),
	}

	return server
}