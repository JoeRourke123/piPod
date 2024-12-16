package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"net/http"
	"orchestrator/service/db"
	"orchestrator/service/spotify"
	"orchestrator/ui/model"
	"orchestrator/util/api"
	"orchestrator/util/logger"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Get(api.LoginRedirect(), redirectToSpotifyLogin)
	app.Get(api.CompleteAuth(), handleAuthComplete)
	app.Get(api.IsAuth(), handleIsAuth)
}

func redirectToSpotifyLogin(ctx *fiber.Ctx) error {
	return ctx.Redirect(spotify.GetAuthUrl())
}

func handleAuthComplete(ctx *fiber.Ctx) error {
	request, _ := adaptor.ConvertRequest(ctx, false)
	token, err := spotify.Auth.Token(ctx.Context(), spotify.AuthState, request)
	if err != nil {
		logger.Error(ctx.Context(), "could not build Spotify Token", err, logger.FromTag("handleAuthComplete"), logger.ApiTag("spotify", "Auth.Token"))
		return ctx.Status(http.StatusInternalServerError).SendString("Sorry could not build Spotify Token")
	}

	tokenResponse, _ := json.Marshal(token)
	logger.Info(ctx.Context(), "token received: "+string(tokenResponse), logger.FromTag("handleAuthComplete"), logger.ApiTag("spotify", "Auth.Token"))

	db.SetSpotifyToken(token)

	return ctx.SendString("Thanks!")
	//return ctx.Redirect("http://pipod.local:3000/login/success")
}

func handleIsAuth(ctx *fiber.Ctx) error {
	token := spotify.GenerateAccessToken(ctx.Context())
	isOnline := db.IsInternetEnabled()

	authResponse := model.AuthResponse{
		HasToken: !isOnline || token != nil,
		AuthUrl:  api.Full(api.LoginRedirect()),
	}

	if !isOnline {
		authResponse.AccessToken = db.GetSpotifyToken().AccessToken
	} else if authResponse.HasToken {
		authResponse.AccessToken = token.AccessToken
	}

	responseJson, _ := json.Marshal(authResponse)

	return ctx.SendString(string(responseJson))
}
