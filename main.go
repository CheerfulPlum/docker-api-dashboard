package main

import (
	"log"
	"net/http"
)

var dockerClient = getDockerClient()

func main() {
	http.HandleFunc("/container/", getContainer)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
