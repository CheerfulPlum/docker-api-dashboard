package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var dockerClient = getDockerClient()

// Initialize
var containerList = listContainers(dockerClient)

// var containerListMutex = sync.Mutex{}

func main() {

	// Update container list async
	go func() {
		for {
			// containerListMutex.Lock()
			fmt.Printf("address of slice %p \n", &containerList[0])
			containerList = listContainers(dockerClient)
			// containerListMutex.Unlock()
			time.Sleep(time.Second)
			fmt.Println("Update")
		}
	}()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
