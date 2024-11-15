package items

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/ui/model"
	"orchestrator/ui/views"
	"orchestrator/util"
)

func PodcastEpisodesToListViewItem(podcastUri spotify.URI, podcastId spotify.ID, episodes []spotify.EpisodePage) []model.ListViewItemResponse {
	return util.Map(episodes, func(e spotify.EpisodePage) model.ListViewItemResponse {
		return model.ListViewItemResponse{
			Title:           e.Name,
			Path:            views.Playing(string(e.URI), string(podcastUri), podcastId.String()),
			BackgroundImage: util.CheckForImage(e.Images),
		}
	})
}

func PodcastsToListViewItem(podcasts []spotify.SavedShow) []model.ListViewItemResponse {
	return util.Map(podcasts, func(pod spotify.SavedShow) model.ListViewItemResponse {
		return model.ListViewItemResponse{
			Title:           pod.Name,
			BackgroundImage: util.CheckForImage(pod.Images),
			Path:            views.Podcast(string(pod.ID)),
		}
	})
}
