package db

import (
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
	"github.com/tomjowitt/gotidal"
	"github.com/zmb3/spotify/v2"
	"time"
)

type DownloadedAlbum struct {
	SpotifyAlbumId   string            `json:"id"`
	DownloadedAt     time.Time         `json:"downloadedAt"`
	TracksToFilePath map[string]string `json:"tracks"`
}

func SetDownloadedAlbum(spotifyAlbum *spotify.FullAlbum, tidalAlbum *gotidal.Album) {
	doc := document.NewDocumentOf(DownloadedAlbum{
		SpotifyAlbumId: spotifyAlbum.ID.String(),
		DownloadedAt:   time.Now(),
	})
	for _, track := range spotifyAlbum.Tracks.Tracks {
		trackFilePath := "./download/" + tidalAlbum.Artists[0].Name + "/" + tidalAlbum.Title + "/" + track.Name + ".flac"
		doc.Set(track.ID.String(), trackFilePath)
	}
	db.UpdateFunc(query.NewQuery(albumCollection).Where(query.Field("album.id").Eq(spotifyAlbum.ID.String())), func(doc *document.Document) *document.Document {
		doc.Set("downloaded", true)
		return doc
	})

	db.InsertOne(downloadsCollection, doc)
}

func IsTrackDownloaded(spotifyTrackId string) bool {
	exists, err := db.Exists(query.NewQuery(downloadsCollection).Where(query.Field(spotifyTrackId).Exists()))
	return err == nil && exists
}

func IsAlbumDownloaded(albumId string) bool {
	exists, _ := db.Exists(query.NewQuery(downloadsCollection).Where(query.Field("SpotifyAlbumId").Eq(albumId)))
	return exists
}

func GetDownloadedAlbums(offset int) []spotify.SavedAlbum {
	albums, _ := db.FindAll(query.NewQuery(albumCollection).Where(query.Field("downloaded").IsTrue()).Limit(15).Skip(offset))
	return mapDocToAlbum(albums)
}
