package spotify

import (
	"conductor/util/logger"
	"context"
	"golang.org/x/oauth2"
)

func Token(ctx context.Context) *oauth2.Token {
	client := GetClient(ctx)
	token, err := client.Token()
	if err != nil {
		logger.Error(
			ctx,
			"error generating new auth token",
			err, logger.ApiTag("spotify", "Token"), logger.FromTag("Token"),
		)
		return nil
	}

	return token
}

func AuthUrl() string {
	return Auth.AuthURL(AuthState)
}
