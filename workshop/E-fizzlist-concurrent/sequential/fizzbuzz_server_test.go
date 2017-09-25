package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"strconv"
	"fmt"
)

func TestApiFizzBuzzRange_1to3_listof3(t *testing.T) {
	server := httptest.NewServer(router()) //starts a real server on a free port
	defer server.Close()                   //notice we use defer here to ensure our server is closed
	req, err := http.NewRequest("POST", server.URL+"/api/fizzbuzzrange", strings.NewReader(`{"from":"1", "to":"3"}`))
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
	message := &MessageResp{}
	if err := json.Unmarshal(echo, message); err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	if (len(message.Results)) != 3 {
		t.Fail()
		log.Println("expected 3 results, got " + strconv.Itoa(len(message.Results)))
	}

	if "1" != message.Results[0].Result {
		t.Fail()
		log.Println("expected the text for number " + strconv.Itoa(message.Results[0].Number) +
			" to be 1, got " + message.Results[0].Result)
	}
	if "2" != message.Results[1].Result {
		t.Fail()
		log.Println("expected the text for number " + strconv.Itoa(message.Results[1].Number) +
			" to be 2, got " + message.Results[1].Result)
	}
	if "fizz" != message.Results[2].Result {
		t.Fail()
		log.Println("expected the text for number " + strconv.Itoa(message.Results[2].Number) +
			" to be 2, got " + message.Results[2].Result)
	}
}
