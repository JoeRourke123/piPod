package spotify

import (
	"context"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
	"orchestrator/service/db"
	"orchestrator/util/logger"
)

func GetAuthUrl() string {
	authUrl := Auth.AuthURL(AuthState)
	logger.Info(context.Background(), "generated auth URL: "+authUrl, logger.FromTag("GetAuthUrl"), logger.ApiTag("spotify", "AuthURL"))
	return authUrl
}

func GenerateAccessToken(ctx context.Context) *oauth2.Token {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))
	token, err := client.Token()
	if err != nil {
		logger.Error(
			ctx,
			"error generating new auth token",
			err, logger.ApiTag("spotify", "Token"), logger.FromTag("GenerateAccessToken"),
		)
	}

	db.SetSpotifyToken(token)

	return token
}
