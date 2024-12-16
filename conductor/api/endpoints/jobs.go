package endpoints

import (
	"conductor/job"
	"conductor/util/api"
	"github.com/gofiber/fiber/v2"
)

var (
	JobsEndpoints = []Endpoints{
		{
			Method:  "GET",
			Path:    api.TriggerJob(":name"),
			Handler: handleTriggerJob,
		},
	}
)

func handleTriggerJob(ctx *fiber.Ctx) error {
	jobName := ctx.Params("name")

	switch jobName {
	case "albums":
		job.RefreshAlbums.Handler(ctx.Context())
	case "os":
		job.UpdateOsState.Handler(ctx.Context())
	case "playlists":
		job.RefreshPlaylists.Handler(ctx.Context())
	}

	return ctx.SendStatus(fiber.StatusOK)
}
