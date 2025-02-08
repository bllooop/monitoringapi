package main

import (
	"time"

	logger "github.com/bllooop/monitoringapi/pkg/logging"

	"github.com/bllooop/monitoringapi/pinger/internal/docker"
	"github.com/bllooop/monitoringapi/pinger/internal/ping"
	"github.com/bllooop/monitoringapi/pinger/internal/utils"
)

func main() {
	for {
		containerIPs := docker.GetContainerIPs("monitoring_network")
		if len(containerIPs) == 0 {
			logger.Log.Error().Msg("No containers found")
		}

		for _, ip := range containerIPs {
			go func(ip string) {
				result := ping.PingContainer(ip)
				utils.SendPingResult(result)
			}(ip)
		}

		time.Sleep(10 * time.Second)
	}
}
