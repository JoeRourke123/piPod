package controller

import (
	"github.com/gofiber/fiber/v2"
	"orchestrator/service/tidal"
	"orchestrator/util/api"
)

func SetupDownloadRoutes(app *fiber.App) {
	app.Get(api.DownloadTrack(":trackId"), handleDownloadTrack)
	app.Get(api.DownloadAlbum(":albumId"), handleDownloadAlbum)
}

func handleDownloadTrack(ctx *fiber.Ctx) error {
	trackId := ctx.Params("trackId")

	go tidal.DownloadTrack(ctx.Context(), trackId)

	return ctx.SendStatus(fiber.StatusOK)
}

func handleDownloadAlbum(ctx *fiber.Ctx) error {
	albumId := ctx.Params("albumId")

	go tidal.DownloadAlbum(ctx.Context(), albumId)

	return ctx.SendStatus(fiber.StatusOK)
}
