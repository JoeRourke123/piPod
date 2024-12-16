package views

import (
	"conductor/common/model"
	"conductor/db/fetch"
	"context"
	"github.com/gofiber/fiber/v2"
)

type SimpleView struct {
	Path           string
	Handler        func(ctx context.Context) model.ListView
	HandlerOffline func(ctx context.Context) model.ListView
}
type viewController struct {
	Endpoints []SimpleView
}

func (vc *viewController) Init(app *fiber.App) {
	for _, endpoint := range vc.Endpoints {
		app.Get(endpoint.Path, func(c *fiber.Ctx) error {
			if fetch.InternetEnabled() {
				return c.JSON(endpoint.Handler(context.Background()))
			} else {
				return c.JSON(endpoint.HandlerOffline(context.Background()))
			}
		})
	}
}

var ViewController = &viewController{
	Endpoints: []SimpleView{
		HomeView,
		MusicView,
	},
}
