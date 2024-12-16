package spotify

import (
	"conductor/common/model"
	"conductor/util"
	"conductor/util/logger"
	"context"
	"github.com/zmb3/spotify/v2"
	"time"
)

func Playlist(ctx context.Context, id string) *model.Playlist {
	client := GetClient(ctx)

	playlist, err := client.GetPlaylist(ctx, spotify.ID(id))
	if err != nil {
		logger.Error(ctx, "error fetching playlist: "+id, err, logger.ApiTag("spotify", "GetPlaylist"), logger.FromTag("Playlist"))
		return nil
	}

	return FullPlaylistParser(playlist)
}

func Playlists(ctx context.Context, offset int) []*model.Playlist {
	client := GetClient(ctx)

	playlists, err := client.CurrentUsersPlaylists(ctx, spotify.Limit(util.MaxAlbumLimit), spotify.Offset(offset))
	if err != nil || playlists == nil {
		if err != nil {
			logger.Error(ctx, "error fetching user playlists", err, logger.ApiTag("spotify", "CurrentUserPlaylists"), logger.FromTag("Playlists"))
		}
		return nil
	}

	notNilPlaylists := util.MapNotNil(playlists.Playlists, func(playlist spotify.SimplePlaylist) *model.Playlist {
		if playlist.ID != "" {
			time.Sleep(time.Millisecond * 100)
			return Playlist(ctx, playlist.ID.String())
		} else {
			return nil
		}
	})

	return util.Point(notNilPlaylists...)
}
