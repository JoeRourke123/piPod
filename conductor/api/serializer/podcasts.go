package serializer

import (
	"conductor/api/builder"
	"conductor/common/model"
	"conductor/util"
	"conductor/util/api"
)

type podcastSerializer struct{}

func (s podcastSerializer) Get(podcast *model.Album) model.ListView {
	return builder.ListView().Title(podcast.Name).ShowStatus(true).Items(TrackSerializer.Items(util.Point(podcast.Tracks...), podcast.Uri)).Build()
}

func (s podcastSerializer) List(podcasts []*model.Album) model.ListView {
	return builder.ListView().Title("Podcasts").ShowStatus(true).Items(s.Items(podcasts)).Build()
}

func (s podcastSerializer) Items(podcasts []*model.Album) []model.ListViewItem {
	return util.Map(podcasts, func(podcast *model.Album) model.ListViewItem {
		return builder.ListViewItem().
			Title(podcast.Name).
			Subtitle(podcast.Artist).
			BackgroundImage(api.Full(api.Artwork(podcast.Id))).
			Path(api.Podcast(podcast.Id)).
			Build()
	})
}

var (
	PodcastSerializer = podcastSerializer{}
)
