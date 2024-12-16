package controller

import (
	"conductor/db/fetch"
	"conductor/util"
	"conductor/util/api"
	"conductor/util/clickwheel"
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/sys/unix"
	"log"
	"time"
)

type socketController struct{}

func (s socketController) init(app *fiber.App) {
	app.Get(api.Websocket(), websocket.New(func(c *websocket.Conn) {
		previousEvent := &clickwheel.ClickWheelEvent{IsClickWheelPressed: false, Button: "ClickWheel", ClickwheelPosition: 0}
		serverFD := util.OpenSocketConnection()

		defer unix.Close(serverFD)

		go func() {
			for {
				if c.Conn == nil {
					break
				}

				osState := fetch.OsState()
				osStateJson, _ := json.Marshal(osState)
				err := c.WriteMessage(websocket.TextMessage, osStateJson)
				if err != nil {
					break
				}
				time.Sleep(10 * time.Second)
			}
		}()

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

	}))
}

var SocketController = socketController{}
