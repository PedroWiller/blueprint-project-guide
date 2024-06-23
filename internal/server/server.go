package server

import (
	"blueprint-project/pkg/helpers"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

type FiberServer struct {
	*fiber.App

	// db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "blueprint-project",
			AppName:      "blueprint-project",
			ErrorHandler: helpers.DefaultErrorHandler,
		}),
		// db: database.New(),
	}

	server.Use(recover.New())

	return server
}
