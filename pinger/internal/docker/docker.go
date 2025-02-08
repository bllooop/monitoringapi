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
	/*var networks []domain.DockerNetwork
	if err := json.Unmarshal(out.Bytes(), &networks); err != nil {
		logger.Log.Error().Err(err).Msg("")
		return nil
	}
	var containerIPs []string
	if len(networks) > 0 {
		for _, container := range networks[0].Containers {
			ip := strings.Split(container.IPv4Address, "/")[0]
			containerIPs = append(containerIPs, ip)
		}
	}

	/*cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		logger.Log.Error().Err(err).Msg("")
		return nil
	}
	containers, err := cli.ContainerList(context.Background(), types.ListOptions{})
	if err != nil {
		logger.Log.Error().Err(err).Msg("")
		return nil
	}
	var containerIPs []string
	for _, container := range containers {
		contJSON, err := cli.ContainerInspect(context.Background(), container.ID)
		if err != nil {
			logger.Log.Error().Err(err).Msg("")
			return nil
		}
		for _, netsetting := range contJSON.NetworkSettings.Networks {
			containerIPs = append(containerIPs, netsetting.IPAddress)
		}
	}
	/*
	   resp, err := http.Get(config.DockerAPIURL)

	   	if err != nil {
	   		logger.Log.Error().Err(err).Msg("")
	   		return nil
	   	}

	   defer resp.Body.Close()

	   body, _ := io.ReadAll(resp.Body)
	   var containers []domain.DockerContainer
	   json.Unmarshal(body, &containers)

	   var containerIPs []string

	   	for _, container := range containers {
	   		for _, network := range container.NetworkSettings.Networks {
	   			if network.IPAddress != "" {
	   				containerIPs = append(containerIPs, network.IPAddress)
	   			}
	   		}
	   	}
	*/
	logger.Log.Debug().Msgf("Discovered container IPs: %v", containerIPs)

	return containerIPs
}
