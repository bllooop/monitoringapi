package main

import (
	"time"

	logger "github.com/bllooop/monitoringapi/pkg/logging"

	"github.com/bllooop/monitoringapi/pinger/internal/docker"
	"github.com/bllooop/monitoringapi/pinger/internal/ping"
)

func main() {
	for {
		containerIPs := docker.GetContainerIPs("monitoring_network")
		if len(containerIPs) == 0 {
			logger.Log.Error().Msg("Контейнеры не найдены")
			return
		}

		for _, ip := range containerIPs {
			go func(ip string) {
				result := ping.PingContainer(ip)
				ping.PingProducer(result)
			}(ip)
		}

		time.Sleep(10 * time.Second)
	}
}
