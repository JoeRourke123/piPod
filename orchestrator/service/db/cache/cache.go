package cache

import (
	"context"
	"github.com/ostafen/clover/v2/document"
	sptfy "github.com/zmb3/spotify/v2"
	"io"
	"net/http"
	"orchestrator/service/db"
	"orchestrator/service/spotify"
	"orchestrator/util"
	"orchestrator/util/logger"
	"os"
	"time"
)

func GetCacheAlbums(offset int) []sptfy.SavedAlbum {
	cachedAlbums, err := getAlbumsFromCache(offset)

	if len(cachedAlbums) > 0 && err == nil {
		return cachedAlbums
	}

	return make([]sptfy.SavedAlbum, 0)
}

func getAlbumsFromCache(offset int) ([]sptfy.SavedAlbum, error) {
	albums, err := db.GetAlbums(offset)
	return albums, err
}

func SaveAlbumsToCache(albums []sptfy.SavedAlbum, offset int) {
	albumPosition := offset
	util.Map(albums, func(album sptfy.SavedAlbum) *document.Document {
		albumPosition += 1
		if len(album.Images) > 0 {
			saveAlbumArtwork(album.ID.String(), album.Images[0].URL)
		}
		return db.InsertAlbum(album, albumPosition)
	})
}

func GetAlbumFromCache(ctx context.Context, albumId string) *sptfy.SavedAlbum {
	cachedAlbum := getAlbumFromCache(albumId)

	if cachedAlbum != nil || !db.IsInternetEnabled() {
		return cachedAlbum
	}

	album := spotify.GetAlbum(ctx, albumId)

	savedAlbum := sptfy.SavedAlbum{AddedAt: time.Now().String(), FullAlbum: *album}
	return &savedAlbum
}

func getAlbumFromCache(albumId string) *sptfy.SavedAlbum {
	album, _ := db.GetAlbum(albumId)
	return album
}

func saveAlbumArtwork(albumId string, artworkUrl string) {
	go func() {
		artworkFilename := "artwork/" + albumId + ".jpeg"
		if _, err := os.Stat(artworkFilename); err == nil {
			return
		}

		f, err := os.Create(artworkFilename)
		res, err := http.Get(artworkUrl)

		if err != nil {
			logger.Error(context.Background(), "could not download artwork: ", err, logger.FromTag("SpotifyBehaviourCache"), logger.ApiTag("spotify", "ArtworkDownload"), logger.DbWriteTag)
		}

		defer res.Body.Close()

		_, err = io.Copy(f, res.Body)
	}()
}
