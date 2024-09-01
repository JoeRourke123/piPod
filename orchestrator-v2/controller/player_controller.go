package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"orchestrator/service/spotify"
	"orchestrator/ui"
)

func SetupPlayerRoutes(app *fiber.App) {
	app.Post("/player", handlePlayer)
}

func handlePlayer(ctx *fiber.Ctx) error {
	var playerRequest ui.PlayerRequest
	err := ctx.BodyParser(&playerRequest)
	if err != nil {
		fmt.Println("error parsing player request:", err)
		return err
	}

	context := ctx.Context()

	fmt.Println(playerRequest)

	switch playerRequest.Action {
	case "START":
		spotify.Start(context, playerRequest.DeviceId, playerRequest.SpotifyUri)
		break
	case "PAUSE":
		spotify.Pause(context, playerRequest.DeviceId)
		break
	case "PLAY":
		spotify.Play(context, playerRequest.DeviceId)
		break
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}
