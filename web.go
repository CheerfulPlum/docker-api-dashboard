package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var (
	regexen = make(map[string]*regexp.Regexp)
	relock  sync.Mutex
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Do some routing here, get the URL path
	var slug string
	responseStatus := 200

	path := r.URL.Path
	// Add some routes in here with regex
	switch {
	case match(path, "/containers/([^/]+)", &slug):
		responseStatus = getContainer(w, r, slug)
	case match(path, "/containers"):
		responseStatus = containersList(w, r)
	default:
		responseStatus = errorResponse(w, http.StatusNotFound, "404: Not found")
	}

	// Log to console
	fmt.Println(time.Now().Format(time.RFC3339) + " - " + fmt.Sprint(responseStatus) + " - " + path)
}

func errorResponse(w http.ResponseWriter, code int, message string) int {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	en := json.NewEncoder(w)
	en.Encode(map[string]string{"message": message})

	return code
}

// Stolen to do regex matching
func match(path, pattern string, vars ...interface{}) bool {
	regex := mustCompileCached(pattern)
	matches := regex.FindStringSubmatch(path)
	if len(matches) <= 0 {
		return false
	}
	for i, match := range matches[1:] {
		switch p := vars[i].(type) {
		case *string:
			*p = match
		case *int:
			n, err := strconv.Atoi(match)
			if err != nil {
				return false
			}
			*p = n
		default:
			return false
		}
	}
	return true
}

// Stolen to do regex matching
func mustCompileCached(pattern string) *regexp.Regexp {
	relock.Lock()
	defer relock.Unlock()

	regex := regexen[pattern]
	if regex == nil {
		regex = regexp.MustCompile("^" + pattern + "$")
		regexen[pattern] = regex
	}
	return regex
}
