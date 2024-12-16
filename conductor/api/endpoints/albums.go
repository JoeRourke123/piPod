package endpoints

import (
	"conductor/data/spotify"
	"conductor/data/tidal"
	"conductor/db/fetch"
	"conductor/db/update"
	"conductor/util/api"
	"github.com/gofiber/fiber/v2"
	"os"
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
		{
			Method:  "POST",
			Path:    api.DownloadAlbum(":id"),
			Handler: handleDownloadAlbum,
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

	artworkFilename := "./.db/artwork/" + albumId + ".jpg"

	_, err := os.Stat(artworkFilename)
	doesNotExist := os.IsNotExist(err)

	if doesNotExist && fetch.InternetEnabled() {
		album := spotify.Album(ctx.Context(), albumId)
		if album == nil {
			podcast := spotify.Podcast(ctx.Context(), albumId)
			if podcast == nil {
				return ctx.SendStatus(fiber.StatusNotFound)
			} else {
				return ctx.Redirect(podcast.CoverArtUrl)
			}

		} else {
			return ctx.Redirect(album.CoverArtUrl)
		}
	}

	return ctx.SendFile(artworkFilename, true)
}

func handleDownloadAlbum(ctx *fiber.Ctx) error {
	albumId := ctx.Params("id")

	album := fetch.Album(albumId)
	isDownloaded := !album.Metadata.IsDownloaded

	if isDownloaded {
		err := tidal.DownloadAlbum(ctx.Context(), albumId)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func pinAlbum(ctx *fiber.Ctx, pin bool) error {
	albumId := ctx.Params("id")

	err := update.Albums.Pinned(albumId, pin)
	if err != nil {
		return ctx.SendStatus(500)
	}

	return ctx.SendStatus(200)
}
