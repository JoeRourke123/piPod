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
	RefreshAlbums = Job{
		Interval: time.Minute * 30,
		Handler: func(ctx context.Context) {
			token := spotify.Token(ctx)
			if token == nil {
				logger.Error(ctx, "could not refresh album cache", errors.New("spotify token not found"))
				return
			}
			err := insert.SpotifyToken(token)
			albums := spotify.Albums(ctx)
			err = insert.Album(albums...)
			artworkErr := saveAlbumArtwork(albums)
			if err != nil {
				logger.Error(ctx, "could not refresh album cache: ", err)
			}
			if artworkErr != nil {
				logger.Error(ctx, "could not save album artwork: ", artworkErr)
			}
			logger.Info(ctx, "refreshed album cache")
		},
	}
)

func saveAlbumArtwork(albums []*model.Album) error {
	for _, album := range albums {
		err := insert.AlbumArtwork(album.Id, album.CoverArtUrl)
		if err != nil {
			return err
		}
	}
	return nil
}
