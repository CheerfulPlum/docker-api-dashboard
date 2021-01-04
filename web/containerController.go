package web

import (
	"docker-dashboard-api/docker"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetContainer GetContainer
func getContainer(w http.ResponseWriter, r *http.Request, slug string) int {
	containerID := slug

	containers := ContainerList
	container, err := docker.FindContainerInList(containerID, containers)

	if err != nil {
		return errorResponse(w, http.StatusNotFound, "Container not found - "+containerID)
	}

	jsonCont, err := json.Marshal(container)
	if err != nil {
		return errorResponse(w, http.StatusInternalServerError, err.Error())
	}

	setHeaders(w)
	fmt.Fprint(w, string(jsonCont))

	return http.StatusOK
}

func getContainersIndex(w http.ResponseWriter, r *http.Request) int {
	jsonCont, err := json.Marshal(ContainerList)
	if err != nil {
		return errorResponse(w, http.StatusInternalServerError, err.Error())
	}

	setHeaders(w)
	fmt.Fprint(w, string(jsonCont))

	return http.StatusOK
}

func getContainerHealth(w http.ResponseWriter, r *http.Request, slug string) int {
	containerID := slug

	containers := ContainerList
	container, err := docker.FindContainerInList(containerID, containers)

	if err != nil {
		return errorResponse(w, http.StatusNotFound, "Container not found - "+containerID)
	}

	json, _ := json.Marshal(docker.ContainerHealth{StatusText: container.Container.Status, IsContainerHealthy: docker.IsContainerHealthy(container)})

	setHeaders(w)
	fmt.Fprint(w, string(json))

	return http.StatusOK
}
