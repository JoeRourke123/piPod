package endpoints

import (
	"conductor/common/model"
	"conductor/data/spotify"
	"conductor/db/fetch"
	"conductor/db/insert"
	"conductor/util/api"
	"conductor/util/logger"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"net/http"
)

var (
	AuthEndpoints = []Endpoints{
		{
			Method:  "GET",
			Path:    api.IsAuth(),
			Handler: handleIsAuth,
		},
		{
			Method:  "GET",
			Path:    api.CompleteAuth(),
			Handler: handleCompleteAuth,
		},
		{
			Method:  "GET",
			Path:    api.LoginRedirect(),
			Handler: redirectToSpotifyLogin,
		},
	}
)

func redirectToSpotifyLogin(ctx *fiber.Ctx) error {
	return ctx.Redirect(spotify.AuthUrl())
}

func handleIsAuth(ctx *fiber.Ctx) error {
	token := fetch.SpotifyToken()
	isOnline := fetch.InternetEnabled()

	authResponse := model.AuthResponse{
		AuthUrl:  api.Full(api.LoginRedirect()),
		HasToken: token != nil,
	}

	if token != nil {
		authResponse.AccessToken = token.AccessToken
	}

	if isOnline {
		freshToken := spotify.Token(ctx.Context())
		if freshToken != nil {
			authResponse.AccessToken = freshToken.AccessToken
		}
	}

	responseJson, _ := json.Marshal(authResponse)

	return ctx.SendString(string(responseJson))
}

func handleCompleteAuth(ctx *fiber.Ctx) error {
	request, _ := adaptor.ConvertRequest(ctx, false)
	token, err := spotify.Auth.Token(ctx.Context(), spotify.AuthState, request)
	if err != nil {
		logger.Error(ctx.Context(), "could not build Spotify Token", err, logger.FromTag("handleAuthComplete"), logger.ApiTag("spotify", "Auth.Token"))
		return ctx.Status(http.StatusInternalServerError).SendString("Sorry could not build Spotify Token")
	}

	tokenResponse, _ := json.Marshal(token)
	logger.Info(ctx.Context(), "token received: "+string(tokenResponse), logger.FromTag("handleAuthComplete"), logger.ApiTag("spotify", "Auth.Token"))

	insert.SpotifyToken(token)

	return ctx.SendString("Thanks!")
}
