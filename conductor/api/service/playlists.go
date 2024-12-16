package service

import (
	"conductor/common/model"
	"conductor/db/fetch"
)

var (
	PlaylistService = Service[model.Playlist]{
		Key:    "playlists",
		Lister: fetch.Playlists,
		Getter: fetch.Playlist,
	}
)
