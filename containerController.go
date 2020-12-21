package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getContainer(w http.ResponseWriter, r *http.Request, slug string) int {
	containerID := slug

	containers := containerList
	container, err := findContainerInList(containerID, containers)

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

func containersList(w http.ResponseWriter, r *http.Request) int {
	// containerListMutex.Lock()
	jsonCont, err := json.Marshal(containerList)
	// containerListMutex.Unlock()
	if err != nil {
		return errorResponse(w, http.StatusInternalServerError, err.Error())
	}

	setHeaders(w)
	fmt.Fprint(w, string(jsonCont))

	return http.StatusOK
}
