package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"orchestrator/service/spotify"
	"orchestrator/ui/responses"
	"orchestrator/util/api"
)

func SetupPodcastsController(app *fiber.App) {
	app.Get(api.PodcastList(), handleGetPodcasts)
	app.Get(api.Podcast(":podcastId"), handleGetPodcast)
}

func handleGetPodcasts(ctx *fiber.Ctx) error {
	offset := ctx.QueryInt("next", 0)
	podcasts := spotify.GetPodcasts(ctx.Context(), offset)

	podcastsResponse := responses.GetPodcastsResponse(podcasts)

	podcastsJson, _ := json.Marshal(podcastsResponse)

	return ctx.Send(podcastsJson)
}

func handleGetPodcast(ctx *fiber.Ctx) error {
	podcastId := ctx.Params("podcastId")
	offset := ctx.QueryInt("next", 0)
	podcast := spotify.GetPodcast(ctx.Context(), podcastId)
	episodes := spotify.GetPodcastEpisodes(ctx.Context(), podcastId, offset)

	podcastResponse := responses.GetPodcastResponse(podcast, episodes)

	podcastJson, _ := json.Marshal(podcastResponse)

	return ctx.Send(podcastJson)
}
