package main

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func getDockerClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		log.Fatal(err)
	}

	return cli
}

func listContainers(client *client.Client) []types.Container {
	containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	return containers
}

func findContainerInList(containerID string, containers []types.Container) (*types.Container, error) {
	for _, container := range containers {
		if container.ID == containerID {
			return &container, nil
		}

	}

	// If we can't find the container return an error
	return nil, errors.New("Container not found")
}

func isContainerHealthy(container *types.Container) bool {
	if strings.Contains(container.Status, "(unhealthy)") {
		return false
	}

	// If container health can't be determined then assume it's okay...
	return true
}
