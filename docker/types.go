package docker

import (
	"github.com/docker/docker/api/types"
)

// ContainerHealth con
type ContainerHealth struct {
	StatusText         string
	IsContainerHealthy bool
}

// Container con
type Container struct {
	Container   types.Container
	VirtualHost string
	VirtualPort int
}
