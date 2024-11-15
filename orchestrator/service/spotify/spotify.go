package spotify

import (
	"context"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"orchestrator/service/db"
	"orchestrator/util"
	"orchestrator/util/logger"
	"os"
)

var (
	spotifyScopes = []string{
		spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopePlaylistReadPrivate, spotifyauth.ScopeUserReadCurrentlyPlaying,
		spotifyauth.ScopeUserModifyPlaybackState, spotifyauth.ScopeUserReadPlaybackState, spotifyauth.ScopeUserLibraryModify,
		spotifyauth.ScopePlaylistModifyPublic, spotifyauth.ScopeStreaming, spotifyauth.ScopePlaylistModifyPrivate,
		spotifyauth.ScopeUserTopRead, spotifyauth.ScopeUserReadRecentlyPlayed, spotifyauth.ScopeUserLibraryRead,
		spotifyauth.ScopeUserFollowRead, spotifyauth.ScopeUserFollowModify, spotifyauth.ScopeUserReadEmail,
	}
	Auth = spotifyauth.New(
		spotifyauth.WithClientID(os.Getenv("SPOTIFY_CLIENT_ID")),
		spotifyauth.WithClientSecret(os.Getenv("SPOTIFY_CLIENT_SECRET")),
		spotifyauth.WithRedirectURL(redirectUrl),
		spotifyauth.WithScopes(spotifyScopes...),
	)
)

func GetAlbums(ctx context.Context, offset int) []spotify.SavedAlbum {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))

	albums, err := client.CurrentUsersAlbums(ctx, spotify.Limit(util.MaxAlbumLimit), spotify.Offset(offset))
	if err != nil {
		logger.Error(
			context.Background(),
			"error fetching user albums",
			err, logger.ApiTag("spotify", "CurrentUserAlbums"), logger.FromTag("GetAlbums"),
		)
	}

	if albums != nil && albums.Albums != nil {
		return albums.Albums
	} else {
		return make([]spotify.SavedAlbum, 0)
	}
}

func GetAlbum(ctx context.Context, albumId string) *spotify.FullAlbum {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))

	album, err := client.GetAlbum(ctx, spotify.ID(albumId))
	if err != nil {
		logger.Error(
			context.Background(),
			"error fetching album: "+albumId,
			err, logger.ApiTag("spotify", "GetAlbum"), logger.FromTag("GetAlbum"),
		)
	}

	return album
}

func GetPlaylists(ctx context.Context, offset int) []spotify.SimplePlaylist {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))

	playlists, err := client.CurrentUsersPlaylists(ctx, spotify.Limit(util.MaxAlbumLimit), spotify.Offset(offset))
	if err != nil {
		logger.Error(
			context.Background(),
			"error fetching user playlists",
			err, logger.ApiTag("spotify", "CurrentUserPlaylists"), logger.FromTag("GetPlaylists"),
		)
	}

	return playlists.Playlists
}

func GetPlaylist(ctx context.Context, playlistId string) *spotify.FullPlaylist {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))

	playlist, err := client.GetPlaylist(ctx, spotify.ID(playlistId))
	if err != nil {
		logger.Error(
			context.Background(),
			"error fetching playlist: "+playlistId,
			err, logger.ApiTag("spotify", "GetPlaylist"), logger.FromTag("GetPlaylist"),
		)
	}

	return playlist
}

func GetPlaylistTracks(ctx context.Context, playlistId string, offset int) []*spotify.FullTrack {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))

	playlistItems, err := client.GetPlaylistItems(context.Background(), spotify.ID(playlistId), spotify.Limit(util.MaxAlbumLimit), spotify.Offset(offset))
	if err != nil {
		logger.Error(
			context.Background(),
			"error fetching playlist tracks: "+playlistId,
			err, logger.ApiTag("spotify", "GetPlaylistItems"), logger.FromTag("GetPlaylistTracks"),
		)
	}

	return util.Map(playlistItems.Items, func(i spotify.PlaylistItem) *spotify.FullTrack {
		return i.Track.Track
	})
}

func GetQueue(ctx context.Context) []spotify.FullTrack {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))

	queue, err := client.GetQueue(ctx)
	if err != nil {
		logger.Error(ctx, "error fetching queue", err, logger.FromTag("GetQueue"), logger.ApiTag("spotify", "GetQueue"))
	}

	return queue.Items
}

func QueueTrack(ctx context.Context, trackId string, deviceId string) {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))

	spotifyDeviceId := spotify.ID(deviceId)
	playOptions := spotify.PlayOptions{DeviceID: &spotifyDeviceId}
	err := client.QueueSongOpt(ctx, spotify.ID(trackId), &playOptions)

	if err != nil {
		logger.Error(ctx, "error queuing track", err, logger.FromTag("QueueTrack"), logger.ApiTag("spotify", "QueueTrack"))
	}
}

func GetPodcasts(ctx context.Context, offset int) []spotify.SavedShow {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))

	showPages, err := client.CurrentUsersShows(ctx, spotify.Limit(util.MaxAlbumLimit), spotify.Offset(offset))
	if err != nil {
		logger.Error(
			context.Background(),
			"error fetching user podcasts",
			err, logger.ApiTag("spotify", "CurrentUsersShows"), logger.FromTag("GetPodcasts"),
		)
		return make([]spotify.SavedShow, 0)
	}

	return showPages.Shows
}

func GetPodcast(ctx context.Context, podcastId string) spotify.FullShow {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))

	podcast, err := client.GetShow(ctx, spotify.ID(podcastId))
	if err != nil {
		logger.Error(
			context.Background(),
			"error fetching podcast: "+podcastId,
			err, logger.ApiTag("spotify", "GetShow"), logger.FromTag("GetPodcast"),
		)
	}

	return *podcast
}

func GetPodcastEpisodes(ctx context.Context, podcastId string, offset int) []spotify.EpisodePage {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))

	episodes, err := client.GetShowEpisodes(ctx, podcastId, spotify.Limit(util.MaxAlbumLimit), spotify.Offset(offset))
	if err != nil {
		logger.Error(
			context.Background(),
			"error fetching podcast episodes: "+podcastId,
			err, logger.ApiTag("spotify", "GetShowEpisodes"), logger.FromTag("GetPodcastEpisodes"),
		)
	}

	return episodes.Episodes
}

func GetEpisode(ctx context.Context, episodeId string) *spotify.EpisodePage {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))

	episode, err := client.GetEpisode(ctx, episodeId)
	if err != nil {
		logger.Error(
			context.Background(),
			"error fetching episode: "+episodeId,
			err, logger.ApiTag("spotify", "GetEpisode"), logger.FromTag("GetEpisode"),
		)
	}

	return episode
}

func GetTrack(ctx context.Context, trackId string) *spotify.FullTrack {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))

	track, err := client.GetTrack(ctx, spotify.ID(trackId))
	if err != nil {
		logger.Error(
			context.Background(),
			"error fetching user tracks",
			err, logger.ApiTag("spotify", "CurrentUsersTracks"), logger.FromTag("GetTracks"),
		)
	}

	return track
}

const (
	redirectUrl = "http://localhost:9091/auth"
	AuthState   = "PIPOD123"
)
