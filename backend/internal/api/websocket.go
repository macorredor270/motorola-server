package api

import (
	"net/http"
	"time"

	"motorola-server/internal/services"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func DeviceSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	defer conn.Close()

	for {

		sysInfo := services.GetSystem()

		data := map[string]interface{}{
			"host":    sysInfo.Host,
			"os":      sysInfo.OS,
			"cpu":     sysInfo.CPU,
			"ram":     sysInfo.RAM,
			"battery": sysInfo.Battery,
			"time":    time.Now(),
		}

		err := conn.WriteJSON(data)
		if err != nil {
			return
		}

		time.Sleep(2 * time.Second)
	}
}
