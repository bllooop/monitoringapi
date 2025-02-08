package utils

import (
	"bytes"
	"encoding/json"
	"net/http"

	config "github.com/bllooop/monitoringapi/pinger/config"
	"github.com/bllooop/monitoringapi/pinger/internal/domain"
	logger "github.com/bllooop/monitoringapi/pkg/logging"
)

func SendPingResult(result domain.PingResult) {
	url := config.BackendURL

	data, err := json.Marshal(result)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to marshal ping result")
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to send ping result")
		return
	}
	defer resp.Body.Close()

	logger.Log.Info().Msgf("Sent ping result, received status: %d", resp.StatusCode)
}
