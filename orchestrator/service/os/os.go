package os

import (
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
	"orchestrator/service/db"
	"time"
)

func SendOsUpdates(c *websocket.Conn) {
	go func() {
		for {
			osUpdates := db.GetOsUpdates()
			osUpdatesJson, _ := json.Marshal(osUpdates)
			c.WriteMessage(websocket.TextMessage, osUpdatesJson)
			time.Sleep(10 * time.Second)
		}
	}()
}
