package ping

import (
	"time"

	config "github.com/bllooop/monitoringapi/pinger/config"
	logger "github.com/bllooop/monitoringapi/pkg/logging"

	domain "github.com/bllooop/monitoringapi/pinger/internal/domain"
	probing "github.com/prometheus-community/pro-bing"
)

func PingContainer(ip string) domain.PingResult {
	pinger, err := probing.NewPinger(ip)
	if err != nil {
		logger.Log.Error().Err(err).Msgf("Error creating pinger %s", ip)
		return domain.PingResult{IP: ip, LastSuccess: domain.LastSuccess}
	}

	pinger.Count = 3
	pinger.Timeout = time.Duration(config.PingTimeout) * time.Second

	err = pinger.Run()
	if err != nil {
		logger.Log.Error().Err(err).Msgf("Ping failed for %s", ip)
		return domain.PingResult{IP: ip, LastSuccess: domain.LastSuccess}
	}

	stats := pinger.Statistics()
	var result domain.PingResult
	if stats.PacketsRecv > 0 {
		domain.LastSuccess = time.Now()
		result = domain.PingResult{
			IP:          ip,
			LastSuccess: domain.LastSuccess,
		}
	} else {
		result = domain.PingResult{
			IP:          ip,
			LastSuccess: domain.LastSuccess,
		}
	}

	logger.Log.Debug().Msgf("Pinged %s - Last_Success: %v", ip, result.LastSuccess)

	return result
}
