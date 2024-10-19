package cache

import (
	"context"
	"github.com/ostafen/clover/v2/document"
	sptfy "github.com/zmb3/spotify/v2"
	"orchestrator/service/db"
	"orchestrator/service/spotify"
	"orchestrator/util"
)

func GetCacheAlbums(ctx context.Context, offset int) []sptfy.SavedAlbum {
	cachedAlbums, err := getAlbumsFromCache(offset)

	if len(cachedAlbums) > 0 && err == nil {
		return cachedAlbums
	}

	db.DeleteAlbums()
	albums := spotify.GetAlbums(ctx, offset)
	saveAlbumsToCache(albums, offset)
	return albums
}

func getAlbumsFromCache(offset int) ([]sptfy.SavedAlbum, error) {
	albums, err := db.GetAlbums(offset)
	return albums, err
}

func saveAlbumsToCache(albums []sptfy.SavedAlbum, offset int) {
	albumPosition := offset
	util.Map(albums, func(album sptfy.SavedAlbum) *document.Document {
		albumPosition += 1
		return db.InsertAlbum(album, albumPosition)
	})
}

func GetAlbumFromCache(ctx context.Context, albumId string) sptfy.FullAlbum {
	cachedAlbum := getAlbumFromCache(albumId)

	if cachedAlbum != nil {
		return *cachedAlbum
	}

	album := spotify.GetAlbum(ctx, albumId)

	return *album
}

func getAlbumFromCache(albumId string) *sptfy.FullAlbum {
	album, _ := db.GetAlbum(albumId)
	return album
}
