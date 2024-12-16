package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"orchestrator/service"
	"orchestrator/service/db"
	"orchestrator/service/db/cache"
	"orchestrator/ui/responses"
	"orchestrator/util/api"
)

func SetupAlbumRoutes(app *fiber.App) {
	app.Get(api.AlbumsList(), getUserAlbums)
	app.Get(api.Album(":albumId"), getAlbum)
	app.Get(api.Artwork(":albumId"), getAlbumArtwork)
	app.Get(api.UnpinAlbum(":albumId"), unpinAlbum)
}

func getAlbum(ctx *fiber.Ctx) error {
	next := ctx.QueryInt("next", 0)
	albumId := ctx.Params("albumId")
	album := cache.GetAlbumFromCache(ctx.Context(), albumId)

	albumResponse := responses.GetAlbumResponse(album)

	if next >= int(album.Tracks.Total) {
		albumResponse = responses.GetEmptyResponse(album.Name)
	}

	albumJson, _ := json.Marshal(albumResponse)

	return ctx.Send(albumJson)
}

func getUserAlbums(ctx *fiber.Ctx) error {
	offset := ctx.QueryInt("next", 0)

	filter := ctx.Query("filter", "")
	sort := ctx.Query("sort", "")
	albums := service.SortFilterAlbums(offset, filter, sort)

	albumResponse := responses.GetAlbumsResponse(albums)

	albumJson, _ := json.Marshal(albumResponse)

	return ctx.Send(albumJson)
}

func getAlbumArtwork(ctx *fiber.Ctx) error {
	albumId := ctx.Params("albumId")

	artworkFilename := "artwork/" + albumId + ".jpeg"

	return ctx.SendFile(artworkFilename, true)
}

func unpinAlbum(ctx *fiber.Ctx) error {
	albumId := ctx.Params("albumId")

	err := db.PinAlbum(albumId, false)
	if err != nil {
		return ctx.SendStatus(500)
	}

	return ctx.SendStatus(200)
}
