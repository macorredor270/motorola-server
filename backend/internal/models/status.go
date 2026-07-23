package models

type DeviceStatus struct {
	Host string `json:"host"`
	OS string `json:"os"`
	CPU float64 `json:"cpu"`
	RAM float64 `json:"ram"`
	Battery int `json:"battery"`
	Status string `json:"status"`
}
