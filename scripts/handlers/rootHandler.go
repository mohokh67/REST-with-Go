package handlers

import (
	"net/http"
)

// RootHandler server the root endpoint
func RootHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found\n"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API v1\n"))
}
