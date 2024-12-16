package job

import (
	"conductor/common/model"
	"conductor/data/spotify"
	"conductor/db/insert"
	"conductor/util/logger"
	"context"
	"errors"
	"time"
)

var (
	RefreshPlaylists = Job{
		Interval: time.Minute * 30,
		Handler: func(ctx context.Context) {
			token := spotify.Token(ctx)
			if token == nil {
				logger.Error(ctx, "could not refresh playlist cache", errors.New("spotify token not found"))
				return
			}
			err := insert.SpotifyToken(token)
			playlists := spotify.Playlists(ctx, 0)
			err = insert.Playlist(playlists...)
			artworkErr := savePlaylistArtwork(playlists)
			if err != nil {
				logger.Error(ctx, "could not refresh playlist cache: ", err)
			}
			if artworkErr != nil {
				logger.Error(ctx, "could not save playlist artwork: ", artworkErr)
			}
			logger.Info(ctx, "refreshed playlist cache")
		},
	}
)

func savePlaylistArtwork(playlists []*model.Playlist) error {
	for _, playlist := range playlists {
		if playlist.CoverArtUrl != "" {
			err := insert.AlbumArtwork(playlist.Id, playlist.CoverArtUrl)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
