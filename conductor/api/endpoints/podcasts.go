package endpoints

import (
	"conductor/data/podcasts"
	"conductor/data/spotify"
	"conductor/db/fetch"
	"conductor/util/api"
	"github.com/gofiber/fiber/v2"
)

var (
	PodcastEndpoints = []Endpoints{
		{
			Method:  "POST",
			Path:    api.DownloadEpisode(":episodeId"),
			Handler: handleDownloadPodcastEpisode,
		},
		{
			Method:  "GET",
			Path:    api.DownloadedPodcastList(),
			Handler: handleDownloadedPodcasts,
		},
	}
)

func handleDownloadPodcastEpisode(ctx *fiber.Ctx) error {
	episodeId := ctx.Params("episodeId", "")
	if episodeId != "" && fetch.InternetEnabled() {
		episode := spotify.Episode(ctx.Context(), episodeId)
		if episode != nil {
			err := podcasts.DownloadEpisode(episode)
			if err == nil {
				return ctx.SendStatus(fiber.StatusOK)
			}
		}
	}

	return ctx.SendStatus(fiber.StatusInternalServerError)
}
