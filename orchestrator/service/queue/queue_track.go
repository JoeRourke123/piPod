package queue

import (
	"context"
	"encoding/json"
	"github.com/ostafen/clover/v2/document"
	sptfy "github.com/zmb3/spotify/v2"
	"orchestrator/service/db"
	"orchestrator/service/db/cache"
	"orchestrator/service/spotify"
	"orchestrator/util"
)

func Track(ctx context.Context, track *sptfy.SimpleTrack, album *sptfy.SimpleAlbum, deviceId string) {
	if db.IsInternetEnabled() && deviceId != "" {
		spotify.QueueTrack(ctx, track.ID.String(), deviceId)
	} else {
		db.QueueTrack(track, album)
	}
}

func Pop(ctx context.Context) (*sptfy.SimpleTrack, *sptfy.SimpleAlbum) {
	if db.IsInternetEnabled() {
		queue := spotify.GetQueue(ctx)
		if len(queue) == 0 {
			return nil, nil
		} else {
			return &queue[0].SimpleTrack, &queue[0].Album
		}
	}
	trackId, albumId := db.PopQueue()
	if trackId == "" || albumId == "" {
		return nil, nil
	}

	track, album, _ := cache.GetTrack(ctx, trackId, albumId)

	return track, &album.SimpleAlbum
}

func List(ctx context.Context) []db.QueueItem {
	if db.IsInternetEnabled() {
		db.ClearQueue()
		return util.Map(spotify.GetQueue(ctx), func(t sptfy.FullTrack) db.QueueItem {
			Track(ctx, &t.SimpleTrack, &t.Album, "")
			return db.QueueItem{Track: &t.SimpleTrack, Album: &t.Album}
		})
	}

	return db.GetQueue()
}

func SetPlaybackContext(ctx context.Context, playbackContext string, spotifyUri string) {
	playbackType := spotify.UriType(playbackContext)

	newPlaybackQueue := make([]*document.Document, 0)

	if playbackType == spotify.AlbumType {
		album, playbackContext := cache.GetAlbumPlaybackContext(ctx, playbackContext, spotifyUri)
		newPlaybackQueue = util.Map(playbackContext, func(track sptfy.SimpleTrack) *document.Document {
			doc := document.NewDocument()
			trackJson, _ := json.Marshal(track)
			albumJson, _ := json.Marshal(album.SimpleAlbum)
			doc.Set("trackId", track.ID.String())
			doc.Set("albumId", album.ID.String())
			doc.Set("track", trackJson)
			doc.Set("album", albumJson)
			return doc
		})
	}

	db.SetQueuePlaybackContext(playbackContext, newPlaybackQueue)
}

func Empty(ctx context.Context) bool {
	return len(List(ctx)) == 0
}
