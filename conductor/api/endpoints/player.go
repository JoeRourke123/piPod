package endpoints

import (
	"conductor/api/serializer"
	"conductor/common/adaptor"
	"conductor/common/model"
	"conductor/db/fetch"
	"conductor/util"
	"conductor/util/api"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

var (
	PlayerEndpoints = []Endpoints{
		{
			Path:   api.Player(),
			Method: "POST",
			Handler: func(ctx *fiber.Ctx) error {
				request := new(model.PlayerRequest)
				if err := ctx.BodyParser(request); err != nil {
					return ctx.SendStatus(fiber.StatusBadRequest)
				}

				track := adaptor.Player(ctx.Context(), request)
				serializedTrack := serializer.PlayerSerializer.Serialize(track)
				jsonTrack, _ := json.Marshal(serializedTrack)
				return ctx.Send(jsonTrack)
			},
		},
		{
			Path:   api.PlayerContent(":uri"),
			Method: "GET",
			Handler: func(ctx *fiber.Ctx) error {
				uri := ctx.Params("uri", "NONE")
				uriType := util.GetTypeFromUri(uri)
				switch uriType {
				case "NONE":
					return ctx.SendStatus(fiber.StatusBadRequest)
				case "TRACK":
					track := fetch.Track(uri)
					trackFileLocation := track.Metadata.FileLocation
					return ctx.SendFile(trackFileLocation, false)
				case "EPISODE":
					episode := fetch.Episode(uri)
					episodeFileLocation := episode.Metadata.FileLocation
					return ctx.SendFile(episodeFileLocation, false)
				}

				return ctx.SendStatus(fiber.StatusBadRequest)
			},
		},
	}
)
