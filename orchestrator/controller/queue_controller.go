package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"orchestrator/service/spotify"
	"orchestrator/ui/responses"
	"orchestrator/util/api"
)

func SetupQueueRoutes(app *fiber.App) {
	app.Get(api.QueueList(), handleGetQueue)
	app.Post(api.QueueTrack(":trackId"), handlePostQueue)
}

func handleGetQueue(ctx *fiber.Ctx) error {
	queue := spotify.GetQueue(ctx.Context())

	queueResponse := responses.GetQueueResponse(queue)

	queueJson, _ := json.Marshal(queueResponse)

	return ctx.Send(queueJson)
}

func handlePostQueue(ctx *fiber.Ctx) error {
	trackId := ctx.Params("trackId")

	type QueueTrackRequest struct {
		DeviceId string `json:"device_id"`
	}
	var request QueueTrackRequest
	ctx.BodyParser(&request)

	//logger.Error(ctx.Context(), "error parsing request", err, logger.FromTag("handlePostQueue"))

	spotify.QueueTrack(ctx.Context(), trackId, request.DeviceId)

	return ctx.SendStatus(fiber.StatusOK)
}
