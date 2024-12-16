package serializer

import (
	"orchestrator/api/model"
	"orchestrator/ui/views"
	"orchestrator/util"
	"orchestrator/util/api"
)

type trackSerializer struct{}

func (s trackSerializer) Get(track *model.Track) model.ListView {
	return model.ListView{}
}

func (s trackSerializer) List(tracks []model.Track) model.ListView {
	return model.ListView{
		Title:      "Tracks",
		ShowStatus: false,
		Items:      s.Items(tracks),
	}
}

func (s trackSerializer) Items(tracks []model.Track) []model.ListViewItem {
	return util.Map(tracks, func(track model.Track) model.ListViewItem {
		return model.ListViewItem{
			Title:           track.Name,
			Path:            views.Playing(track.Uri, track.Album.Id, track.Album.Id),
			BackgroundImage: api.Artwork(track.Album.Id),
		}
	})
}

var (
	TrackSerializer = trackSerializer{}
)
