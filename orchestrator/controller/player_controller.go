package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"orchestrator/service/db"
	"orchestrator/service/db/cache"
	"orchestrator/service/player"
	"orchestrator/service/spotify"
	"orchestrator/ui/model"
	"orchestrator/ui/responses"
	"orchestrator/util"
	"orchestrator/util/api"
	"strings"
)

func SetupPlayerRoutes(app *fiber.App) {
	app.Post(api.Player(), handlePlayer)
	app.Get(api.PlayerContent(":contentId"), handlePlayerContent)
}

func handlePlayer(ctx *fiber.Ctx) error {
	var playerRequest model.PlayerRequest
	err := ctx.BodyParser(&playerRequest)
	if err != nil {
		fmt.Println("error parsing player request:", err)
		return err
	}

	playerService := player.PlayerServiceBuilder(playerRequest)

	track, album := playerService(ctx.Context(), playerRequest)

	if track != nil && album != nil {
		db.SetCurrentTrack(track, album, playerRequest.PlaybackContext)
		response := responses.GetTrackPlayerResponse(track, album)
		responseJson, _ := json.Marshal(response)
		return ctx.Send(responseJson)
	} else if strings.Contains(playerRequest.SpotifyUri, "track") {
		track, album, _ := cache.GetTrack(ctx.Context(), playerRequest.SpotifyUri, playerRequest.AlbumID)
		db.SetCurrentTrack(track, &album.SimpleAlbum, playerRequest.PlaybackContext)
		response := responses.GetTrackPlayerResponse(track, &album.SimpleAlbum)
		responseJson, _ := json.Marshal(response)
		return ctx.Send(responseJson)
	} else if strings.Contains(playerRequest.SpotifyUri, "show") {
		podcast := spotify.GetPodcast(ctx.Context(), playerRequest.AlbumID)
		episode := spotify.GetEpisode(ctx.Context(), util.UriToId(playerRequest.SpotifyUri))
		response := responses.GetPodcastPlayerResponse(podcast, episode)
		responseJson, _ := json.Marshal(response)
		return ctx.Send(responseJson)
	} else {
		response := responses.GetCurrentPlayerResponse()
		responseJson, _ := json.Marshal(response)
		return ctx.Send(responseJson)
	}
}

func handlePlayerContent(ctx *fiber.Ctx) error {
	contentId := ctx.Params("contentId")

	filePath := db.GetTrackPath(contentId)

	if filePath != "" {
		return ctx.SendFile(filePath, false)
	} else {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
}
