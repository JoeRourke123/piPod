package serializer

import (
	"conductor/api/builder"
	"conductor/common/constants"
	"conductor/common/model"
	"conductor/db/fetch"
	"conductor/util/api"
)

type queueSerializer struct{}

// Get do not use! /*
func (s queueSerializer) Get(_ *model.Track) model.ListView {
	return builder.ListView().Build()
}

func (s queueSerializer) List(queue []*model.Track) model.ListView {
	return builder.ListView().Title("Queue").ShowStatus(true).Items(s.Items(queue)).Build()
}

func (s queueSerializer) Items(queue []*model.Track) []model.ListViewItem {
	if len(queue) == 0 {
		return make([]model.ListViewItem, 0)
	}

	nowPlaying := queue[0]

	playbackContextUri := fetch.PlaybackContext()

	nowPlayingItem := builder.ListViewItem().
		Title(nowPlaying.Name).
		Path("/playing/" + nowPlaying.Uri + "/" + playbackContextUri).
		Icon(constants.SPEAKER_HIGH).
		BackgroundImage(api.Full(api.Artwork(nowPlaying.Album.Id))).
		Build()

	queuedItems := TrackSerializer.Items(queue[1:], playbackContextUri)

	return append([]model.ListViewItem{nowPlayingItem}, queuedItems...)
}

var (
	QueueSerializer = queueSerializer{}
)
