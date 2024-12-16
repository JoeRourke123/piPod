package endpoints

import "github.com/gofiber/fiber/v2"

type Endpoints struct {
	Path    string
	Method  string
	Handler func(*fiber.Ctx) error
}
