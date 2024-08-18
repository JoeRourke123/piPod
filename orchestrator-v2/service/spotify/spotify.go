package spotify

import (
	"context"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"log"
	"orchestrator/service/db"
	"orchestrator/util"
	"os"
)

var (
	spotifyScopes = []string{
		spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopePlaylistReadPrivate, spotifyauth.ScopeUserReadCurrentlyPlaying,
		spotifyauth.ScopeUserModifyPlaybackState, spotifyauth.ScopeUserReadPlaybackState, spotifyauth.ScopeUserLibraryModify,
		spotifyauth.ScopePlaylistModifyPublic, spotifyauth.ScopeStreaming, spotifyauth.ScopePlaylistModifyPrivate,
		spotifyauth.ScopeUserTopRead, spotifyauth.ScopeUserReadRecentlyPlayed, spotifyauth.ScopeUserLibraryRead,
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
		log.Println("error fetching user albums: ", err)
	}

	return albums.Albums
}

func GetAlbum(ctx context.Context, albumId string) *spotify.FullAlbum {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))

	album, err := client.GetAlbum(ctx, spotify.ID(albumId))
	if err != nil {
		log.Println("error fetching album: ", albumId, err)
	}

	return album
}

func GetPlaylists(ctx context.Context, offset int) []spotify.SimplePlaylist {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))

	playlists, err := client.CurrentUsersPlaylists(ctx, spotify.Limit(util.MaxAlbumLimit), spotify.Offset(offset))
	if err != nil {
		log.Println("error fetching user playlists: ", err)
	}

	return playlists.Playlists
}

func GetPlaylist(ctx context.Context, playlistId string) *spotify.FullPlaylist {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))

	playlist, err := client.GetPlaylist(ctx, spotify.ID(playlistId))
	if err != nil {
		log.Println("error fetching playlist: ", playlistId, err)
	}

	return playlist
}

func GetPlaylistTracks(ctx context.Context, playlistId string, offset int) []*spotify.FullTrack {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))

	playlistItems, err := client.GetPlaylistItems(context.Background(), spotify.ID(playlistId), spotify.Limit(util.MaxAlbumLimit), spotify.Offset(offset))
	if err != nil {
		log.Println("error fetching playlist items: ", playlistId, err)
	}

	return util.Map(playlistItems.Items, func(i spotify.PlaylistItem) *spotify.FullTrack {
		return i.Track.Track
	})
}

const (
	redirectUrl = "http://localhost:9091/auth"
	AuthState   = "PIPOD123"
)
