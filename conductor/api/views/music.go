package views

import (
	"conductor/api/builder"
	"conductor/api/serializer"
	"conductor/common/constants"
	"conductor/common/model"
	"conductor/db/fetch"
	"conductor/util/api"
	"context"
)

var (
	MusicView = SimpleView{
		Path: "/music",
		Handler: func(ctx context.Context) model.ListView {
			return builder.ListView().
				Title("Music").
				ShowStatus(true).
				Item(builder.ListViewItem().Title("Albums").Icon(constants.MUSIC_NOTES_SIMPLE).Path(api.AlbumsList()).Build()).
				Item(builder.ListViewItem().Title("Playlists").Icon(constants.MICROPHONE_STAGE).Path(api.PlaylistList()).Build()).
				Item(builder.ListViewItem().Title("Made For You").Icon(constants.CASSETTE_TAPE).Path(api.MadeForYou()).Build()).
				Items(serializer.AlbumSerializer.Items(fetch.PinnedAlbums())).
				Build()
		},
		HandlerOffline: func(ctx context.Context) model.ListView {
			return builder.ListView().Title("Music").ShowStatus(true).
				Item(builder.ListViewItem().Title("Albums").Icon(constants.MUSIC_NOTES_SIMPLE).Path(api.DownloadedAlbumsList()).Build()).
				Item(builder.ListViewItem().Title("Playlists").Icon(constants.MICROPHONE_STAGE).Path(api.DownloadedPlaylistsList()).Build()).
				Items(serializer.AlbumSerializer.Items(fetch.PinnedDownloadedAlbums())).
				Build()
		},
	}
)
