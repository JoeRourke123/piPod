package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"orchestrator/ui/responses"
)

func SetupGamesRoutes(app *fiber.App) {
	app.Get("/list/games", handleListGames)
}

func handleListGames(ctx *fiber.Ctx) error {
	gamesResponse := responses.GetGamesResponse()
	gamesJson, _ := json.Marshal(gamesResponse)
	return ctx.Send(gamesJson)
}
