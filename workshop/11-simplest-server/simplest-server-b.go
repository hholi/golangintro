package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}

	server.ListenAndServe()
	// Accepts connection but does nothing
}
