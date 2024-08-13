package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"golang.org/x/sys/unix"
	"log"
	"orchestrator/clickwheel"
	"orchestrator/controller"
	"orchestrator/service/db"

	"github.com/gofiber/contrib/websocket"
)

func main() {
	db.InitialiseDatabase()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	controller.SetupAlbumRoutes(app)
	controller.SetupAuthRoutes(app)

	defer db.CloseDatabases()

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		previousEvent := &clickwheel.ClickWheelEvent{IsClickWheelPressed: false, Button: "ClickWheel", ClickwheelPosition: 0}
		serverFD := openSocketConnection()

		defer unix.Close(serverFD)

		for {
			response := make([]byte, PacketSize)

			_, _, err := unix.Recvfrom(serverFD, response, 0)
			if err == nil {
				event := clickwheel.BuildClickWheelEvent(previousEvent, int(response[0]), int(response[1]), int(response[2]))
				log.Println(event)
				eventJson, _ := json.Marshal(event)

				if err := c.WriteMessage(websocket.TextMessage, eventJson); err != nil {
					log.Println("write:", err)
					break
				}

				previousEvent = event
			}
		}

	}))

	log.Fatal(app.Listen("0.0.0.0:9091"))
	// Access the websocket server: ws://localhost:3000/ws/123?v=1.0
	// https://www.websocket.org/echo.html
}
