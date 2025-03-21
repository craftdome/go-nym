package models

type Health struct {
	Status string `json:"status"`

	// Uptime in seconds
	Uptime uint64 `json:"uptime"`
}
