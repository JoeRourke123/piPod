package db

import (
	"errors"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
	"github.com/zmb3/spotify/v2"
	"time"
)

func GetTrack(trackUri string, albumId string) (*spotify.SimpleTrack, *spotify.SavedAlbum, error) {
	album, err := GetAlbum(albumId)

	if err != nil {
		return nil, nil, nil
	}
	if album == nil {
		return nil, nil, errors.New("could not find album saved locally")
	}

	for _, track := range album.Tracks.Tracks {
		if trackUri == string(track.URI) || trackUri == string(track.ID) {
			return &track, album, nil
		}
	}

	return nil, album, errors.New("could not find track saved locally")
}

func GetTrackPath(trackId string) string {
	doc, err := db.FindFirst(query.NewQuery(downloadsCollection).Where(query.Field(trackId).Exists()))
	if doc != nil && err == nil {
		return doc.Get(trackId).(string)
	}
	return ""
}

func GetCurrentTrackId() string {
	doc, err := db.FindFirst(query.NewQuery(historyCollection).Sort(query.SortOption{Field: "playedAt", Direction: -1}))
	if doc != nil && err == nil {
		return doc.Get("trackId").(string)
	}

	return ""
}

func GetCurrentAlbumId() string {
	doc, err := db.FindFirst(query.NewQuery(historyCollection).Sort(query.SortOption{Field: "playedAt", Direction: -1}))
	if doc != nil && err == nil {
		return doc.Get("albumId").(string)
	}

	return ""
}

func GetCurrentPlaybackContext() string {
	doc, err := db.FindFirst(query.NewQuery(historyCollection).Sort(query.SortOption{Field: "playedAt", Direction: -1}))
	if doc != nil && err == nil {
		return doc.Get("playbackContext").(string)
	}

	return ""
}

func GetCurrentTrack() (*spotify.SimpleTrack, *spotify.SavedAlbum, string) {
	trackId, albumId, playbackContext := GetCurrentTrackId(), GetCurrentAlbumId(), GetCurrentPlaybackContext()

	if trackId == "" || albumId == "" {
		return nil, nil, ""
	}

	track, album, err := GetTrack(trackId, albumId)
	if err != nil {
		return nil, nil, ""
	}

	return track, album, playbackContext
}

func SetCurrentTrack(track *spotify.SimpleTrack, album *spotify.SimpleAlbum, playbackContext string) {
	doc := document.NewDocument()
	doc.Set("albumId", album.ID.String())
	doc.Set("trackId", track.ID.String())
	doc.Set("playbackContext", playbackContext)
	doc.Set("playedAt", time.Now())
	db.InsertOne(historyCollection, doc)
}
