package main

import (
	"docker-dashboard-api/docker"
	"docker-dashboard-api/web"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// Refresh interval in seconds
var refreshInterval = 5

var defaultLogLevel = "info"

func main() {
	godotenv.Load()
	file := setupLogging()
	// Defer in the main function so we don't prematurely close
	defer file.Close()
	setupRefreshInterval()

	// Update container list async
	go func() {
		for {
			// Refresh the list of containerss
			docker.RefreshContainerList()
			web.ContainerList = docker.ContainerList
			log.Debug("Updating container list")
			log.Debug("Success, " + strconv.Itoa(len(docker.ContainerList)) + " containers available")
			log.Trace(docker.ContainerList)
			// Refresh after a period of time
			time.Sleep(time.Duration(refreshInterval) * time.Second)

		}
	}()

	http.HandleFunc("/", web.Handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func setupLogging() *os.File {
	setupLoggingLevel()
	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.SetOutput(io.MultiWriter(file, os.Stdout))

	return file
}

func setupRefreshInterval() {
	if len(os.Getenv("REFRESH_INTERVAL")) > 0 {
		converted, err := strconv.Atoi(os.Getenv("REFRESH_INTERVAL"))

		if err != nil {
			log.Error(err)
		}

		if err != nil || converted < 1 {
			log.Warn("REFRESH_INTERVAL must be at least 1, defaulting to " + strconv.Itoa(refreshInterval))
			return
		}

		// If we get here we're good to set it to the var
		refreshInterval = converted
	}
}

func setupLoggingLevel() {
	level := os.Getenv("LOGGING_LEVEL")

	if level == "" {
		log.Warn("No LOGGING LEVEL set, setting to default: " + defaultLogLevel)
		parsedLevel, _ := log.ParseLevel(defaultLogLevel)
		log.SetLevel(parsedLevel)
		return
	}

	// If we get here we're good to set it to the var
	parsedLevel, _ := log.ParseLevel(level)

	log.SetLevel(parsedLevel)
}
