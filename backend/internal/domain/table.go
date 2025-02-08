package domain

import "time"

type PingResult struct {
	IP          string    `json:"ip"`
	LastSuccess time.Time `json:"last_success"`
	Timestamp   time.Time `json:"timestamp"`
}
