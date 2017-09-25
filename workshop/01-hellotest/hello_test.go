package main

import (
	"log"
	"testing"
)

func TestHello(t *testing.T) {
	result := Hello("Hallgeir")
	if "Hello Hallgeir!" != result {
		t.Fail()
		log.Println("expected a greeting")
	}
}
