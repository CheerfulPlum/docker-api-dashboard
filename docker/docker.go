package docker

import (
	"context"
	"errors"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	// ContainerList con
	ContainerList []Container
	dockerClient  *client.Client
)

func init() {
	var err error
	dockerClient, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		log.Fatal(err)
	}
}

// ListContainers get an an array of containers back
func ListContainers() []Container {
	containers, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	conList := make([]Container, len(containers))

	for i := 0; i < len(containers); i++ {
		virutalHost, virtualPort := GetContainerVirutalEnvs(containers[i].ID)
		conList[i] = Container{Container: containers[i], VirtualHost: virutalHost, VirtualPort: virtualPort}
	}

	return conList
}

// FindContainerInList find a container by it's ID in a list of contaienrs
func FindContainerInList(containerID string, containers []Container) (*Container, error) {
	for _, container := range containers {
		if container.Container.ID == containerID {
			return &container, nil
		}

	}

	// If we can't find the container return an error
	return nil, errors.New("Container not found")
}

// IsContainerHealthy return a bool based on whether the container is healthy or not
func IsContainerHealthy(container *Container) bool {
	if strings.Contains(container.Container.Status, "(unhealthy)") {
		return false
	}

	// If container health can't be determined then assume it's okay...
	return true
}

// GetContainerVirutalEnvs return the virtual host and port env var from the container
func GetContainerVirutalEnvs(containerID string) (string, int) {
	vHostString := "VIRTUAL_HOST="
	vPortString := "VIRTUAL_PORT="

	// Grab the env vars for the container
	inspect, _ := dockerClient.ContainerInspect(context.Background(), containerID)
	envs := inspect.Config.Env
	vHost := ""
	vPort := 0
	for i := 0; i < len(envs); i++ {
		if strings.Contains(envs[i], vHostString) {
			vHost = envs[i][len(vHostString):]
		} else if strings.Contains(envs[i], vPortString) {
			vPort, _ = strconv.Atoi(envs[i][len(vPortString):])
		}

		if vHost != "" && vPort != 0 {
			break
		}
	}

	return vHost, vPort
}
