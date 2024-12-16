package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"orchestrator/service/db"
	"orchestrator/util/api"
)

var (
	AlbumEndpoints = []Endpoints{
		{
			Method:  "POST",
			Path:    api.PinAlbum(":id"),
			Handler: handlePinAlbum,
		},
		{
			Method:  "POST",
			Path:    api.UnpinAlbum(":id"),
			Handler: handleUnpinAlbum,
		},
		{
			Method:  "GET",
			Path:    api.Artwork(":id"),
			Handler: handleGetArtwork,
		},
	}
)

func handlePinAlbum(ctx *fiber.Ctx) error {
	return pinAlbum(ctx, true)
}

func handleUnpinAlbum(ctx *fiber.Ctx) error {
	return pinAlbum(ctx, false)
}

func handleGetArtwork(ctx *fiber.Ctx) error {
	albumId := ctx.Params("id")

	artworkFilename := "artwork/" + albumId + ".jpeg"

	return ctx.SendFile(artworkFilename, true)
}

func pinAlbum(ctx *fiber.Ctx, pin bool) error {
	albumId := ctx.Params("albumId")

	err := db.PinAlbum(albumId, pin)
	if err != nil {
		return ctx.SendStatus(500)
	}

	return ctx.SendStatus(200)
}
