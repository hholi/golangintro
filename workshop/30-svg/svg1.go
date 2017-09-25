package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ajstarks/svgo"
)

func main() {
	http.Handle("/circle", http.HandlerFunc(circle))
	http.Handle("/circles", http.HandlerFunc(circles))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func circle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(250, 250, 125, "fill:none;stroke:black")
	s.End()
}

func circles(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	for radius := 125; radius > 10; radius-- {
		s.Circle(250, 250, radius, "fill:none;stroke:black")
		time.Sleep(1 * time.Second)
	}
	s.End()
}
