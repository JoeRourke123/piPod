package spotify

import (
	"conductor/db/fetch"
	"context"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"os"
)

var (
	//redirectUrl   = api.Full(api.CompleteAuth())
	redirectUrl   = "http://localhost:9091/auth"
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
	AuthState = "PIPOD123"
)

func GetClient(ctx context.Context) *spotify.Client {
	return spotify.New(Auth.Client(ctx, fetch.SpotifyToken()))
}
