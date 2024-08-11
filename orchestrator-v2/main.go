package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"

	"github.com/gofiber/contrib/websocket"
)

func main() {
	app := fiber.New()

	eventsChannel := make(chan *ClickWheelEvent)

	go openSocketConnection(eventsChannel)

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		for {
			event := <-eventsChannel
			fmt.Println(event)
			eventJson, _ := json.Marshal(event)
			fmt.Println(eventJson)

			if err := c.WriteMessage(websocket.TextMessage, eventJson); err != nil {
				log.Println("write:", err)
				break
			}
		}

	}))

	log.Fatal(app.Listen("0.0.0.0:9091"))
	// Access the websocket server: ws://localhost:3000/ws/123?v=1.0
	// https://www.websocket.org/echo.html
}
