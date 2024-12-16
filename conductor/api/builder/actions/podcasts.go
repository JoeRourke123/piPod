package actions

import (
	"conductor/api/builder"
	"conductor/common/constants"
	"conductor/common/model"
	"conductor/util/api"
)

func DownloadEpisode(episode *model.Track) model.ListViewItem {
	if episode.Metadata.IsDownloaded {
		return builder.ListViewItem().
			Title("Remove Download").
			Icon(constants.X_CIRCLE).
			ActionType("POST").
			RequestUrl(api.Full(api.RemoveDownloadEpisode(episode.Id))).Build()
	} else {
		return builder.ListViewItem().
			Title("Download Episode").
			Icon(constants.DOWNLOAD_SIMPLE).
			ActionType("POST").
			RequestUrl(api.Full(api.DownloadEpisode(episode.Id))).Build()
	}
}

func GoToPodcast(episode *model.Track) model.ListViewItem {
	return builder.ListViewItem().
		Title("Go to Podcast").
		Icon(constants.MICROPHONE_STAGE).
		ActionType("REDIRECT").
		Path(api.Podcast(episode.Album.Id)).Build()
}

func DownloadedPodcasts() model.ListViewItem {
	return builder.ListViewItem().
		Title("Downloaded Podcasts").
		Icon(constants.DOWNLOAD_SIMPLE).
		ActionType("REDIRECT").
		Path(api.DownloadedPodcastList()).Build()
}
