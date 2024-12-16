package cache

import (
	"context"
	"errors"
	"github.com/zmb3/spotify/v2"
	"orchestrator/service/db"
	sptfy "orchestrator/service/spotify"
	"orchestrator/util"
)

func GetTrack(ctx context.Context, trackUriId string, albumId string) (*spotify.SimpleTrack, *spotify.SavedAlbum, error) {
	album := GetAlbumFromCache(ctx, albumId)

	if album != nil {
		for _, track := range album.Tracks.Tracks {
			if trackUriId == string(track.URI) || trackUriId == string(track.ID) {
				return &track, album, nil
			}
		}
	}

	return nil, album, errors.New("could not find track saved locally")
}

func GetAlbumPlaybackContext(ctx context.Context, playbackContext string, spotifyUri string) (*spotify.SavedAlbum, []spotify.SimpleTrack) {
	album := GetAlbumFromCache(ctx, util.UriToId(playbackContext))

	if album != nil {
		for i, track := range album.Tracks.Tracks {
			if spotifyUri == string(track.URI) {
				return album, album.Tracks.Tracks[i+1:]
			}
		}
	}

	return nil, make([]spotify.SimpleTrack, 0)
}

func GetCurrentTrack(ctx context.Context) (*spotify.SimpleTrack, *spotify.SimpleAlbum, *spotify.URI) {
	if db.IsInternetEnabled() {
		_, track, playbackContext := sptfy.IsCurrentlyPlaying(ctx)
		if track == nil {
			return nil, nil, nil
		}

		return &track.SimpleTrack, &track.Album, playbackContext
	} else {
		track, album, playbackContext := db.GetCurrentTrack()
		playbackContextUri := spotify.URI(playbackContext)

		if track == nil || album == nil {
			return nil, nil, nil
		}

		return track, album, &playbackContextUri
	}
}
