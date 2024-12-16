package serializer

import (
	"conductor/api/builder"
	"conductor/api/builder/actions"
	"conductor/common/constants"
	"conductor/common/model"
	"conductor/util"
	"conductor/util/api"
)

type albumSerializer struct{}

func (s albumSerializer) Get(album *model.Album) model.ListView {
	return builder.ListView().
		Title(album.Name).
		ShowStatus(true).
		Items(TrackSerializer.Items(util.Point(album.Tracks...), album.Uri)).
		AdditionalInfo(s.Info(album)...).
		Build()
}

func (s albumSerializer) List(albums []*model.Album) model.ListView {
	return builder.ListView().
		Title("Albums").
		ShowStatus(true).
		Items(s.Items(albums)).
		Build()
}

func (s albumSerializer) Items(albums []*model.Album) []model.ListViewItem {
	return util.Map(albums, func(album *model.Album) model.ListViewItem {
		return builder.ListViewItem().
			Title(album.Name).
			Subtitle(album.Artist).
			BackgroundImage(api.Full(api.Artwork(album.Id))).
			Path("/albums/" + album.Id).
			Action(actions.QueueAlbum(album)).
			Action(actions.DownloadAlbum(album)).
			Action(actions.PinAlbum(album)).
			Build()
	})
}

func (s albumSerializer) Info(album *model.Album) []model.ListViewInfo {
	return []model.ListViewInfo{
		builder.ListViewInfo().Text(album.Artist).Icon(constants.USER_SOUND).Build(),
		builder.ListViewInfo().Text(album.ReleaseDate).Icon(constants.CALENDAR).Build(),
		builder.ListViewInfo().Text(album.Duration).Icon(constants.HOURGLASS_MEDIUM).Build(),
	}
}

var (
	AlbumSerializer = albumSerializer{}
)
