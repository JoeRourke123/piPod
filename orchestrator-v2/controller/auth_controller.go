package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"net/http"
	"orchestrator/service/db"
	"orchestrator/service/spotify"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Get("/go", redirectToSpotifyLogin)
	app.Get("/auth", handleAuthComplete)
}

func redirectToSpotifyLogin(ctx *fiber.Ctx) error {
	return ctx.Redirect(spotify.GetAuthUrl())
}

func handleAuthComplete(ctx *fiber.Ctx) error {
	request, _ := adaptor.ConvertRequest(ctx, false)
	token, err := spotify.Auth.Token(ctx.Context(), spotify.AuthState, request)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("Sorry could not build Spotify Token")
	}

	db.SetSpotifyToken(token)

	return ctx.SendString("Thanks!")
	//return ctx.Redirect("http://pipod.local:3000/login/success")
}
