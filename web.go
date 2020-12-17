package main

import (
	"fmt"
	"net/http"
)

func getContainer(w http.ResponseWriter, r *http.Request) {

	containerID := r.URL.Path[len("/container/"):]

	containers := listContainers(dockerClient)
	container, err := findContainerInList(containerID, containers)

	if err != nil {
		fmt.Fprintf(w, "Cannot find container %s!", containerID)
		return
	}

	health := "Healthy"
	if !isContainerHealthy(container) {
		health = "Unhealthy"
	}

	fmt.Fprintf(w, "Status for container - Container is %s", health)
}
