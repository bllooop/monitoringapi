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
		logger.Log.Error().Err(err).Msgf("Ошибка создания pinger %s", ip)
		return domain.PingResult{IP: ip, LastSuccess: domain.LastSuccess, Timestamp: time.Now()}
	}

	pinger.Count = 3
	pinger.Timeout = time.Duration(config.PingTimeout) * time.Second

	err = pinger.Run()
	if err != nil {
		logger.Log.Error().Err(err).Msgf("Ошибка пинга для %s", ip)
		return domain.PingResult{IP: ip, LastSuccess: domain.LastSuccess, Timestamp: time.Now()}
	}

	stats := pinger.Statistics()
	var result domain.PingResult
	if stats.PacketsRecv > 0 {
		domain.LastSuccess = time.Now()
		result = domain.PingResult{
			IP:          ip,
			LastSuccess: domain.LastSuccess,
			Timestamp:   time.Now(),
		}
	} else {
		result = domain.PingResult{
			IP:          ip,
			LastSuccess: domain.LastSuccess,
			Timestamp:   time.Now(),
		}
	}

	logger.Log.Debug().Msgf("Пинг %s - Успешная попытка: %v", ip, result.LastSuccess)

	return result
}
