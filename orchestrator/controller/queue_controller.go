package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"orchestrator/service/db/cache"
	"orchestrator/service/queue"
	"orchestrator/ui/responses"
	"orchestrator/util/api"
)

func SetupQueueRoutes(app *fiber.App) {
	app.Get(api.QueueList(), handleGetQueue)
	app.Post(api.QueueTrack(":trackId"), handlePostQueue)
}

func handleGetQueue(ctx *fiber.Ctx) error {
	queueItems := queue.List(ctx.Context())

	queueResponse := responses.GetQueueResponse(ctx.Context(), queueItems)

	queueJson, _ := json.Marshal(queueResponse)

	return ctx.Send(queueJson)
}

func handlePostQueue(ctx *fiber.Ctx) error {
	trackId := ctx.Params("trackId")

	type QueueTrackRequest struct {
		DeviceId string `json:"device_id"`
		AlbumId  string `json:"album_id"`
	}
	var request QueueTrackRequest
	ctx.BodyParser(&request)

	//logger.Error(ctx.Context(), "error parsing request", err, logger.FromTag("handlePostQueue"))
	track, album, err := cache.GetTrack(ctx.Context(), trackId, request.AlbumId)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	queue.Track(ctx.Context(), track, &album.SimpleAlbum, request.DeviceId)

	return ctx.SendStatus(fiber.StatusOK)
}
