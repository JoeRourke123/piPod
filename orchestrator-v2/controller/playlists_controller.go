package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"orchestrator/service/spotify"
	"orchestrator/ui"
)

func SetupPlaylistRoutes(app *fiber.App) {
	app.Get("/list/playlists", getUserPlaylists)
	app.Get("/playlists/:playlistId", getPlaylist)
}

func getPlaylist(ctx *fiber.Ctx) error {
	playlistId := ctx.Params("playlistId")
	offset := ctx.QueryInt("next", 0)
	playlist := spotify.GetPlaylist(ctx.Context(), playlistId)
	playlistTracks := spotify.GetPlaylistTracks(ctx.Context(), playlistId, offset)

	playlistResponse := ui.GetPlaylistResponse(playlist, playlistTracks)

	playlistJson, _ := json.Marshal(playlistResponse)

	return ctx.Send(playlistJson)
}

func getUserPlaylists(ctx *fiber.Ctx) error {
	offset := ctx.QueryInt("next", 0)
	playlists := spotify.GetPlaylists(ctx.Context(), offset)

	playlistResponse := ui.GetPlaylistsResponse(playlists)

	playlistJson, _ := json.Marshal(playlistResponse)

	return ctx.Send(playlistJson)
}
