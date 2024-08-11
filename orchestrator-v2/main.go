package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/sys/unix"
	"log"

	"github.com/gofiber/contrib/websocket"
)

func main() {
	app := fiber.New()

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		previousEvent := &ClickWheelEvent{IsClickWheelPressed: false, Button: "ClickWheel", ClickwheelPosition: 0}
		serverFD := openSocketConnection()

		defer unix.Close(serverFD)

		for {
			response := make([]byte, PacketSize)

			_, _, err := unix.Recvfrom(serverFD, response, 0)
			if err == nil {
				event := BuildClickWheelEvent(previousEvent, int(response[0]), int(response[1]), int(response[2]))
				log.Println(event)
				eventJson, _ := json.Marshal(event)

				if err := c.WriteMessage(websocket.TextMessage, eventJson); err != nil {
					log.Println("write:", err)
					break
				}

				previousEvent = event
			} else {
				log.Println("error reading msg: ", err)
			}
		}

	}))

	log.Fatal(app.Listen("0.0.0.0:9091"))
	// Access the websocket server: ws://localhost:3000/ws/123?v=1.0
	// https://www.websocket.org/echo.html
}
