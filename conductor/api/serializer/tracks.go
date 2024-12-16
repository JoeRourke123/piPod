package serializer

import (
	"conductor/api/builder"
	"conductor/api/builder/actions"
	"conductor/common/model"
	"conductor/util"
	"conductor/util/api"
)

type trackSerializer struct{}

func (s trackSerializer) Get(_ *model.Track) model.ListView {
	return builder.ListView().Build()
}

func (s trackSerializer) List(tracks []*model.Track) model.ListView {
	return builder.ListView().Title("Tracks").ShowStatus(true).Items(s.Items(tracks, "")).Build()
}

func (s trackSerializer) Items(tracks []*model.Track, playbackContext string) []model.ListViewItem {
	return util.Map(tracks, func(track *model.Track) model.ListViewItem {
		playbackContext := playbackContext
		if playbackContext == "" {
			playbackContext = track.Uri
		}
		return builder.ListViewItem().
			Title(track.Name).
			Subtitle(track.Artist).
			BackgroundImage(api.Full(api.Artwork(track.Album.Id))).
			Path("/playing/" + track.Uri + "/" + playbackContext).
			Action(actions.QueueTrack(track)).
			Action(actions.Download(track)).
			Action(actions.GoTo(track)).
			Build()
	})
}

var (
	TrackSerializer = trackSerializer{}
)
