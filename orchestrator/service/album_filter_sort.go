package service

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/service/db"
	"orchestrator/service/db/cache"
)

const (
	DOWNLOAD_FILTER = "downloaded"
)

func SortFilterAlbums(offset int, filter string, sort string) []spotify.SavedAlbum {
	albums := make([]spotify.SavedAlbum, 0)

	if filter == DOWNLOAD_FILTER {
		albums = db.GetDownloadedAlbums(offset)
	} else {
		albums = cache.GetCacheAlbums(offset)
	}
	
	return albums
}
