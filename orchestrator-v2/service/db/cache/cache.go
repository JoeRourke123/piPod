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
	cachedAlbums := getAlbumsFromCache(offset)

	if len(cachedAlbums) > 0 {
		return cachedAlbums
	}

	albums := spotify.GetAlbums(ctx, offset)
	saveAlbumsToCache(albums)
	return albums
}

func getAlbumsFromCache(offset int) []sptfy.SavedAlbum {
	albums, _ := db.GetAlbums(offset)
	return albums
}

func saveAlbumsToCache(albums []sptfy.SavedAlbum) {
	util.Map(albums, func(album sptfy.SavedAlbum) *document.Document {
		return db.InsertAlbum(album.FullAlbum)
	})
}

func GetAlbumFromCache(ctx context.Context, albumId string) sptfy.FullAlbum {
	cachedAlbum := getAlbumFromCache(albumId)

	if cachedAlbum != nil {
		return *cachedAlbum
	}

	album := spotify.GetAlbum(ctx, albumId)
	db.InsertAlbum(*album)
	return *album
}

func getAlbumFromCache(albumId string) *sptfy.FullAlbum {
	album, _ := db.GetAlbum(albumId)
	return album
}
