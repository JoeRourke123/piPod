package service

import (
	"conductor/common/model"
	"conductor/db/fetch"
)

var (
	AlbumService = Service[model.Album]{
		Key:    "albums",
		Lister: fetch.Albums,
		Getter: fetch.Album,
	}
)
