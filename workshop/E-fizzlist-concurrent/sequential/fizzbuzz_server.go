package main

// To run:
// 	go run fizzbuzz_server.go fizzbuzz.go
// 	curl -XPOST http://localhost:3001/api/fizzbuzzrange -H 'Content-type:application/json' -d '{"from":"1", "to":"16"}'

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"fmt"
)

//MessageReq wraps a request message and stamps it
type MessageReq struct {
	From 	string `json:"from"` //tells the decoder what to decode into
	To 		string `json:"to"` //tells the decoder what to decode into
	Stamp  	int64  `json:"stamp,omitempty"`
}

//MessageResp wraps a request message and stamps it
type MessageResp struct {
	Results []FizzBuzzResult `json:"fizzbuzzresult"`
	Stamp  int64  `json:"stamp,omitempty"`
}

//BusinessLogic does awesome BuisnessLogic
func FizzBuzzRange(from string, to string) *MessageResp {
	mess := &MessageResp{}
	if from, err := strconv.Atoi(from); err == nil {
		if to, err := strconv.Atoi(to); err == nil {
			mess.Results = FizzBuzzFromTo(from, to)
			mess.Stamp = time.Now().Unix()
		}
	}
	return mess
}

//FizzBuzzRangeHttp return result for a range
func FizzBuzzRangeHttp(res http.ResponseWriter, req *http.Request) {
	var (
		jsonDecoder = json.NewDecoder(req.Body) //decoder reading from the post body
		jsonEncoder = json.NewEncoder(res)      //encoder writing to the response stream
		request     = &MessageReq{}             // something to hold our data
	)
	res.Header().Add("Content-type", "application/json")
	if err := jsonDecoder.Decode(request); err != nil { //decode our data into our struct
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := FizzBuzzRange(request.From, request.To)
	fmt.Println(result)
	if err := jsonEncoder.Encode(result); err != nil { //encode our data and write it back to the response stream
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//Setup our simple router
func router() http.Handler {
	http.HandleFunc("/api/fizzbuzzrange", FizzBuzzRangeHttp)
	return http.DefaultServeMux //this is a stdlib http.Handler
}

func main() {
	router := router()
	//start our server on port 3001
	if err := http.ListenAndServe(":3001", router); err != nil {
		log.Fatal(err)
	}
}
