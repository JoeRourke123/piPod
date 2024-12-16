package controller

import (
	"github.com/gofiber/fiber/v2"
	"orchestrator/api/endpoints"
	"orchestrator/api/serializer"
	"orchestrator/api/service"
)

type Controller[T any] struct {
	Service         service.Service[T]
	Serializer      serializer.Serializer[T]
	CustomEndpoints []endpoints.Endpoints
}

func InitControllers(app *fiber.App) {
	AlbumController.init(app)
}

func (c Controller[T]) init(app *fiber.App) {
	app.Get("/"+c.Service.Key+"/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		item := c.Service.Get(id)
		return ctx.JSON(c.Serializer.Get(item))
	})
	app.Get("/"+c.Service.Key, func(ctx *fiber.Ctx) error {
		items := c.Service.List()
		return ctx.JSON(c.Serializer.List(items))
	})
	for _, endpoint := range c.CustomEndpoints {
		if endpoint.Method == "GET" {
			app.Get(endpoint.Path, endpoint.Handler)
		} else if endpoint.Method == "POST" {
			app.Post(endpoint.Path, endpoint.Handler)
		} else if endpoint.Method == "PUT" {
			app.Put(endpoint.Path, endpoint.Handler)
		} else if endpoint.Method == "DELETE" {
			app.Delete(endpoint.Path, endpoint.Handler)
		}
	}
}
