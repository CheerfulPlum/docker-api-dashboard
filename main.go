package main

import (
	"log"
	"net/http"
)

var dockerClient = getDockerClient()

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
