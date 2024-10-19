package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"orchestrator/service/db/cache"
	"orchestrator/ui/responses"
	"orchestrator/util/api"
)

func SetupAlbumRoutes(app *fiber.App) {
	app.Get(api.AlbumsList(), getUserAlbums)
	app.Get(api.Album(":albumId"), getAlbum)
}

func getAlbum(ctx *fiber.Ctx) error {
	next := ctx.QueryInt("next", 0)
	albumId := ctx.Params("albumId")
	album := cache.GetAlbumFromCache(ctx.Context(), albumId)

	albumResponse := responses.GetAlbumResponse(&album)

	if next >= int(album.Tracks.Total) {
		albumResponse = responses.GetEmptyResponse(album.Name)
	}

	albumJson, _ := json.Marshal(albumResponse)

	return ctx.Send(albumJson)
}

func getUserAlbums(ctx *fiber.Ctx) error {
	offset := ctx.QueryInt("next", 0)

	albums := cache.GetCacheAlbums(ctx.Context(), offset)

	albumResponse := responses.GetAlbumsResponse(albums)

	albumJson, _ := json.Marshal(albumResponse)

	return ctx.Send(albumJson)
}
