package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)

	s.App.Get("/health", s.healthHandler)
	s.App.Get("/error", s.errorHandler)
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString("<h1>Welcome to blueprint project guide!</h1>")
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "UP",
	}

	return c.JSON(resp)
}

func (s *FiberServer) errorHandler(c *fiber.Ctx) error {
	return fiber.NewError(fiber.ErrInternalServerError.Code, "what's wrong with you?")
}
