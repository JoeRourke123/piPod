package responses

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/ui"
	"orchestrator/ui/items"
	"orchestrator/ui/model"
)

func GetPodcastsResponse(podcasts []spotify.SavedShow) model.ListViewResponse {
	return model.ListViewResponse{
		Title:      "Podcasts",
		Items:      items.PodcastsToListViewItem(podcasts),
		Icon:       ui.MICROPHONE_STAGE,
		ShowStatus: true,
	}
}

func GetPodcastResponse(podcast spotify.FullShow, episodes []spotify.EpisodePage) model.ListViewResponse {
	return model.ListViewResponse{
		Title:      podcast.Name,
		Items:      items.PodcastEpisodesToListViewItem(podcast.URI, podcast.ID, episodes),
		ShowStatus: false,
	}
}
