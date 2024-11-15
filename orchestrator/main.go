package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"orchestrator/behaviour"
	"orchestrator/controller"
	"orchestrator/service/db"
	"orchestrator/util/logger"
)

func main() {
	db.InitialiseDatabase()
	app := fiber.New()

	behaviour.Setup()

	logger.Info(context.Background(), "starting PiPod server", "service")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	controller.SetupAlbumRoutes(app)
	controller.SetupAuthRoutes(app)
	controller.SetupPlaylistRoutes(app)
	controller.SetupPlayerRoutes(app)
	controller.SetupQueueRoutes(app)
	controller.SetupPodcastsController(app)
	controller.SetupDownloadRoutes(app)
	controller.SetupDbRoutes(app)
	controller.SetupViewsRoutes(app)
	controller.SetupGamesRoutes(app)

	defer db.CloseDatabases()

	controller.SetupWebsocketRoute(app)

	log.Fatal(app.Listen("0.0.0.0:9091"))
}
