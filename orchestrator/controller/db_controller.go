package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"orchestrator/service/db"
	"orchestrator/util/api"
)

func SetupDbRoutes(app *fiber.App) {
	app.Get(api.Collections(), handleGetCollections)
	app.Get(api.Collection(":collectionName"), handleGetCollectionContent)
	app.Delete(api.Collection(":collectionName"), handleClearCollection)
}

func handleClearCollection(ctx *fiber.Ctx) error {
	err := db.ClearCollection(ctx.Params("collectionName"))

	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func handleGetCollectionContent(ctx *fiber.Ctx) error {
	collectionContent := db.GetCollectionContent(ctx.Params("collectionName"))

	if collectionContent == "" {
		return ctx.SendStatus(fiber.StatusNoContent)
	}

	return ctx.SendString(collectionContent)
}

func handleGetCollections(ctx *fiber.Ctx) error {
	collections := db.GetCollections()
	jsonResponse, err := json.Marshal(collections)

	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Send(jsonResponse)
}
