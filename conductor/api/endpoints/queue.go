package endpoints

import (
	"conductor/common/adaptor"
	"conductor/util/api"
	"github.com/gofiber/fiber/v2"
)

type QueueRequest struct {
	Uri string `json:"uri"`
}

var (
	QueueEndpoints = []Endpoints{
		{
			Method: "POST",
			Path:   api.Queue(":uri"),
			Handler: func(ctx *fiber.Ctx) error {
				uri := ctx.Params("uri", "NONE")

				if uri == "NONE" {
					return ctx.SendStatus(fiber.StatusBadRequest)
				}

				err := adaptor.Queue(ctx.Context(), uri)
				if err != nil {
					return ctx.SendStatus(fiber.StatusInternalServerError)
				}

				return ctx.SendStatus(fiber.StatusOK)
			},
		},
	}
)
