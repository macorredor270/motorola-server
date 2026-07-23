package handlers

import (
	"encoding/json"
	"motorola-server/internal/services"
	"net/http"
)

func Status(w http.ResponseWriter, r *http.Request) {

	sysInfo := services.GetSystem()

	json.NewEncoder(w).Encode(map[string]interface{}{
		"host":    sysInfo.Host,
		"os":      sysInfo.OS,
		"cpu":     sysInfo.CPU,
		"ram":     sysInfo.RAM,
		"battery": sysInfo.Battery,
		"status":  "online",
	})
}
