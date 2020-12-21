package docker

import (
	"context"
	"errors"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	// ContainerList con
	ContainerList []types.Container
)

// GetDockerClient get a new docker client
func GetDockerClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		log.Fatal(err)
	}

	return cli
}

// ListContainers get an an array of containers back
func ListContainers(client *client.Client) []types.Container {
	containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	return containers
}

// FindContainerInList find a container by it's ID in a list of contaienrs
func FindContainerInList(containerID string, containers []types.Container) (*types.Container, error) {
	for _, container := range containers {
		if container.ID == containerID {
			return &container, nil
		}

	}

	// If we can't find the container return an error
	return nil, errors.New("Container not found")
}

// IsContainerHealthy return a bool based on whether the container is healthy or not
func IsContainerHealthy(container *types.Container) bool {
	if strings.Contains(container.Status, "(unhealthy)") {
		return false
	}

	// If container health can't be determined then assume it's okay...
	return true
}
