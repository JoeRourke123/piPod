package endpoints

import (
	"conductor/data/tidal"
	"conductor/db/fetch"
	"conductor/util/api"
	"github.com/gofiber/fiber/v2"
)

var (
	PlaylistEndpoints = []Endpoints{
		{
			Method:  "POST",
			Path:    api.DownloadPlaylist(":playlistId"),
			Handler: handleDownloadPlaylist,
		},
	}
)

func handleDownloadPlaylist(ctx *fiber.Ctx) error {
	playlistId := ctx.Params("playlistId", "")
	if playlistId != "" && fetch.InternetEnabled() {
		err := tidal.DownloadPlaylist(ctx.Context(), playlistId)
		if err == nil {
			return ctx.SendStatus(fiber.StatusOK)
		}
	}

	return ctx.SendStatus(fiber.StatusInternalServerError)
}
