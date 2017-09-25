package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s", request.URL.Path[1])
}

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	server.ListenAndServe()
	//
}
