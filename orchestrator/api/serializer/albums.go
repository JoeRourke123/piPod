package serializer

import (
	"orchestrator/api/model"
	"orchestrator/util"
)

type albumSerializer struct{}

func (s albumSerializer) Get(album *model.Album) model.ListView {
	return model.ListView{
		Title:      album.Name,
		ShowStatus: true,
		Items:      TrackSerializer.Items(album.Tracks),
	}
}

func (s albumSerializer) List(albums []*model.Album) model.ListView {
	return model.ListView{
		Title:      "Albums",
		ShowStatus: false,
		Items:      s.Items(albums),
	}
}

func (s albumSerializer) Items(albums []*model.Album) []model.ListViewItem {
	return util.Map(albums, func(album *model.Album) model.ListViewItem {
		return model.ListViewItem{
			Title: album.Name,
			Path:  "/albums/" + album.Id,
		}
	})
}

var (
	AlbumSerializer = albumSerializer{}
)
