package controller

import (
	"conductor/api/endpoints"
	"conductor/api/serializer"
	"conductor/api/service"
	"conductor/api/views"
	"github.com/gofiber/fiber/v2"
)

type Controller[T any] struct {
	Service         *service.Service[T]
	Serializer      serializer.Serializer[T]
	CustomEndpoints []endpoints.Endpoints
}

func Init(app *fiber.App) {
	AlbumController.init(app)
	PlaylistController.init(app)
	PodcastController.init(app)
	QueueController.init(app)
	SocketController.init(app)
	views.ViewController.Init(app)
	control(endpoints.AuthEndpoints).init(app)
	control(endpoints.PlayerEndpoints).init(app)
	control(endpoints.DatabaseEndpoints).init(app)
	control(endpoints.JobsEndpoints).init(app)
	control(endpoints.DownloadedEndpoints).init(app)
}

func (c Controller[T]) init(app *fiber.App) {
	if c.Service != nil && c.Serializer != nil {
		app.Get("/"+c.Service.Key+"/:id", func(ctx *fiber.Ctx) error {
			id := ctx.Params("id")
			item := c.Service.Get(id)
			if item == nil {
				return ctx.SendStatus(fiber.StatusNotFound)
			} else {
				return ctx.JSON(c.Serializer.Get(item))
			}
		})
		app.Get("/"+c.Service.Key, func(ctx *fiber.Ctx) error {
			offset := ctx.QueryInt("offset", 0)
			items := c.Service.List(offset)
			return ctx.JSON(c.Serializer.List(items))
		})
	}
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

func control(endpoints []endpoints.Endpoints) Controller[any] {
	return Controller[any]{
		CustomEndpoints: endpoints,
	}
}
