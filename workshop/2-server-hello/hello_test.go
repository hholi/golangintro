package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	server := httptest.NewServer(router()) //starts a real server on a free port
	defer server.Close()                   //notice we use defer here to ensure our server is closed
	req, err := http.NewRequest("POST", server.URL+"/api/echo", strings.NewReader(`{"message":"Hei", "name":"Hallgeir"}`))
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	res, err := client.Do(req) // make the call to our test server
	if err != nil {
		log.Fatal(err)
	}

	echo, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	message := &Message{}
	if err := json.Unmarshal(echo, message); err != nil {
		log.Fatal(err)
	}
	if "Hei Hallgeir" != message.Message {
		t.Fail()
		log.Println("expected the message to equal Hei Hallgeir, received " + message.Message)
	}
}
