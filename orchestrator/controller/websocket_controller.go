package controller

import (
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/sys/unix"
	"log"
	"orchestrator/clickwheel"
	"orchestrator/util"
	"orchestrator/util/api"
)

func SetupWebsocketRoute(app *fiber.App) {
	app.Get(api.Websocket(), websocketController())
}

func websocketController() fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		previousEvent := &clickwheel.ClickWheelEvent{IsClickWheelPressed: false, Button: "ClickWheel", ClickwheelPosition: 0}
		serverFD := util.OpenSocketConnection()

		defer unix.Close(serverFD)

		for {
			response := make([]byte, util.PacketSize)

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

	})
}
