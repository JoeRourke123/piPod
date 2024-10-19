package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"orchestrator/ui/model"
	"orchestrator/ui/responses"
	"orchestrator/util/api"
)

func SetupViewsRoutes(app *fiber.App) {
	app.Get(api.HomeView(), handleView(responses.GetHomeResponse))
	app.Get(api.MusicView(), handleView(responses.GetMusicResponse))
}

func handleView(response func() model.ListViewResponse) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		responseJson, _ := json.Marshal(response())
		return ctx.Send(responseJson)
	}
}
