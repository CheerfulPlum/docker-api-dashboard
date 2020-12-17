package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getContainer(w http.ResponseWriter, r *http.Request, slug string) int {
	containerID := slug

	containers := listContainers(dockerClient)
	container, err := findContainerInList(containerID, containers)

	if err != nil {
		return errorResponse(w, http.StatusNotFound, "Container not found - "+containerID)
	}

	jsonCont, err := json.Marshal(container)
	if err != nil {
		return errorResponse(w, http.StatusInternalServerError, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonCont))

	return http.StatusOK
}

func containersList(w http.ResponseWriter, r *http.Request) int {

	containers := listContainers(dockerClient)

	jsonCont, err := json.Marshal(containers)
	if err != nil {
		return errorResponse(w, http.StatusInternalServerError, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonCont))

	return http.StatusOK
}
