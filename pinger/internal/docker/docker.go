package docker

import (
	"bytes"
	"os/exec"
	"strings"

	logger "github.com/bllooop/monitoringapi/pkg/logging"
)

func GetContainerIPs(networkName string) []string {
	//cmd := exec.Command("docker", "network", "inspect", networkName)
	cmd := exec.Command("bash", "-c", "docker inspect --format '{{.Name}} - {{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker ps -q)")

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		logger.Log.Error().Err(err).Msg(stderr.String())
		return nil
	}

	lines := strings.Split(out.String(), "\n")
	var containerIPs []string
	for _, line := range lines {
		if len(line) > 0 {
			parts := strings.Split(line, " - ")
			if len(parts) == 2 {
				containerIPs = append(containerIPs, parts[1])
			}
		}
	}

	logger.Log.Debug().Msgf("IP контейнеров: %v", containerIPs)

	return containerIPs
}
