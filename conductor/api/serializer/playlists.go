package serializer

import (
	"conductor/api/builder"
	"conductor/api/builder/actions"
	"conductor/common/constants"
	"conductor/common/model"
	"conductor/util"
	"conductor/util/api"
)

type playlistSerializer struct{}

func (s playlistSerializer) Get(playlist *model.Playlist) model.ListView {
	return builder.ListView().
		Title(playlist.Name).
		AdditionalInfo(s.Info(playlist)...).
		ShowStatus(false).
		Items(TrackSerializer.Items(util.Point(playlist.Tracks...), playlist.Uri)).
		Build()
}

func (s playlistSerializer) List(playlists []*model.Playlist) model.ListView {
	return builder.ListView().
		Title("Playlists").
		ShowStatus(true).
		Items(s.Items(playlists)).
		Build()
}

func (s playlistSerializer) Items(playlists []*model.Playlist) []model.ListViewItem {
	return util.Map(playlists, func(playlist *model.Playlist) model.ListViewItem {
		return builder.ListViewItem().
			Title(playlist.Name).
			Path(api.Playlist(playlist.Id)).
			BackgroundImage(api.Full(api.Artwork(playlist.Id))).
			Action(actions.DownloadPlaylist(playlist)).
			Build()
	})
}

func (s playlistSerializer) Info(playlist *model.Playlist) []model.ListViewInfo {
	return []model.ListViewInfo{
		builder.ListViewInfo().Icon(constants.HOURGLASS_MEDIUM).Text(playlist.Duration).Build(),
		builder.ListViewInfo().Icon(constants.USER_SOUND).Text(playlist.Owner).Build(),
	}
}

var (
	PlaylistSerializer = playlistSerializer{}
)
