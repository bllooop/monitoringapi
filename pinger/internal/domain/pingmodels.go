package domain

import "time"

type DockerContainer struct {
	ID              string `json:"Id"`
	Names           []string
	NetworkSettings struct {
		Networks map[string]struct {
			IPAddress string `json:"IPAddress"`
		} `json:"Networks"`
	} `json:"NetworkSettings"`
}

type PingResult struct {
	IP          string    `json:"ip"`
	LastSuccess time.Time `json:"last_success"`
	Timestamp   time.Time `json:"timestamp"`
}

type DockerNetwork struct {
	Containers map[string]struct {
		IPv4Address string `json:"IPv4Address"`
	} `json:"Containers"`
}

var LastSuccess time.Time
