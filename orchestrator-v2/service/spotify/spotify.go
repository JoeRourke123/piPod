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

	albums, err := client.CurrentUsersAlbums(context.Background(), spotify.Limit(util.MaxAlbumLimit), spotify.Offset(offset))
	if err != nil {
		log.Println("error fetching user albums: ", err)
	}

	return albums.Albums
}

const (
	redirectUrl = "http://localhost:9091/auth"
	AuthState   = "PIPOD123"
)
