package db

import (
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
	"github.com/tomjowitt/gotidal"
	"github.com/zmb3/spotify/v2"
)

func SetDownloadedTrack(spotifyTrackId string, spotifyAlbumId string, track gotidal.Track) {
	SetDownloadedTrackWithPath(spotifyTrackId, spotifyAlbumId, "./download/"+track.Artists[0].Name+"/"+track.Album.Title+"/"+track.Title+".flac")
}

func SetDownloadedTrackWithPath(spotifyTrackId string, spotifyAlbumId string, trackPath string) {
	doc := document.NewDocument()
	doc.Set("id", spotifyTrackId)
	doc.Set("albumId", spotifyAlbumId)
	doc.Set("path", trackPath)
	db.InsertOne(downloadsCollection, doc)
}

func SetDownloadedAlbum(spotifyAlbum *spotify.FullAlbum, tidalAlbum *gotidal.Album) {
	for _, track := range spotifyAlbum.Tracks.Tracks {
		SetDownloadedTrackWithPath(string(track.ID), string(spotifyAlbum.ID), "./download/"+tidalAlbum.Artists[0].Name+"/"+tidalAlbum.Title+"/"+track.Name+".flac")
	}
}

func IsTrackDownloaded(spotifyTrackId string) bool {
	exists, err := db.Exists(query.NewQuery(downloadsCollection).Where(query.Field("id").Eq(spotifyTrackId)))
	return err == nil && exists
}

func IsAlbumDowwnloaded(albumId string, expectedTrackCount int) bool {
	actualTrackCount, err := db.Count(query.NewQuery(downloadsCollection).Where(query.Field("albumId").Eq(albumId)))
	return err == nil && expectedTrackCount == actualTrackCount
}
