package endpoints

import "github.com/gofiber/fiber/v2"

type Endpoints struct {
	Path    string
	Method  string
	Handler func(*fiber.Ctx) error
}

var (
	FOUR_OH_FOUR = func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
)
