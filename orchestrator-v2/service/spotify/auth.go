package spotify

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
	"orchestrator/service/db"
)

func GetAuthUrl() string {
	return Auth.AuthURL(AuthState)
}

func GenerateAccessToken(ctx context.Context) *oauth2.Token {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))
	token, err := client.Token()
	if err != nil {
		fmt.Println("error generating new token: ", err)
	}

	db.SetSpotifyToken(token)

	return token
}
