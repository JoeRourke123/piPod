package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"orchestrator/service/db/cache"
	"orchestrator/ui"
)

func SetupAlbumRoutes(app *fiber.App) {
	app.Get("/list/albums", getUserAlbums)
	app.Get("/albums/:albumId", getAlbum)
}

func getAlbum(ctx *fiber.Ctx) error {
	albumId := ctx.Params("albumId")
	album := cache.GetAlbumFromCache(ctx.Context(), albumId)

	albumResponse := ui.GetAlbumResponse(&album)

	albumJson, _ := json.Marshal(albumResponse)

	return ctx.Send(albumJson)
}

func getUserAlbums(ctx *fiber.Ctx) error {
	offset := ctx.QueryInt("next", 0)

	albums := cache.GetCacheAlbums(ctx.Context(), offset)

	albumResponse := ui.GetAlbumsResponse(albums)

	albumJson, _ := json.Marshal(albumResponse)

	return ctx.Send(albumJson)
}
