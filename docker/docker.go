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
func ListContainers(client *client.Client) []Container {
	containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	conList := make([]Container, len(containers))

	for i := 0; i < len(containers); i++ {
		conList[i] = Container{Container: containers[i], VirtualHost: GetContainerVirutalHost(containers[i].ID, client), VirtualPort: GetContainerVirutalPort(containers[i].ID, client)}
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

// GetContainerVirutalHost return a string indicating the Virtual host environment variable for the container
func GetContainerVirutalHost(containerID string, client *client.Client) string {
	vHostString := "VIRTUAL_HOST="

	// Grab the env vars for the container
	inspect, _ := client.ContainerInspect(context.Background(), containerID)
	envs := inspect.Config.Env
	vHost := ""
	for i := 0; i < len(envs); i++ {
		if strings.Contains(envs[i], vHostString) {
			vHost = envs[i][len(vHostString):]
			break
		}
	}

	return vHost
}

// GetContainerVirutalPort return an int indicating the Virtual port environment variable for the container
func GetContainerVirutalPort(containerID string, client *client.Client) int {
	vHostString := "VIRTUAL_PORT="

	// Grab the env vars for the container
	inspect, _ := client.ContainerInspect(context.Background(), containerID)
	envs := inspect.Config.Env
	vPort := 0
	for i := 0; i < len(envs); i++ {
		if strings.Contains(envs[i], vHostString) {
			vPort, _ = strconv.Atoi(envs[i][len(vHostString):])
			break
		}
	}

	return vPort
}
