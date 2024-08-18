package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"orchestrator/controller"
	"orchestrator/service/db"
)

func main() {
	db.InitialiseDatabase()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	controller.SetupAlbumRoutes(app)
	controller.SetupAuthRoutes(app)
	controller.SetupPlaylistRoutes(app)

	defer db.CloseDatabases()

	app.Get("/ws", controller.WebsocketController())

	log.Fatal(app.Listen("0.0.0.0:9091"))
	// Access the websocket server: ws://localhost:3000/ws/123?v=1.0
	// https://www.websocket.org/echo.html
}
