package main

// To run:
// 	go run server.go
//	curl -XPOST http://localhost:3001/api/echo -H 'Content-type:application/json' -d '{"number":"15"}'

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

//MessageReq wraps a request message and stamps it
type MessageReq struct {
	Number string `json:"number"` //tells the decoder what to decode into
	Stamp  int64  `json:"stamp,omitempty"`
}

//MessageResp wraps a request message and stamps it
type MessageResp struct {
	Result string `json:"result"` //tells the decoder what to decode into
	Stamp  int64  `json:"stamp,omitempty"`
}

//BusinessLogic does awesome BuisnessLogic
func BusinessLogic(text string) *MessageResp {
	mess := &MessageResp{}
	if num, err := strconv.Atoi(text); err == nil {
		mess.Result = FizzBuzz(num)
		mess.Stamp = time.Now().Unix()
	}
	return mess
}

//Echo echoes what you send
func Echo(res http.ResponseWriter, req *http.Request) {
	var (
		jsonDecoder = json.NewDecoder(req.Body) //decoder reading from the post body
		jsonEncoder = json.NewEncoder(res)      //encoder writing to the response stream
		message     = &MessageReq{}             // something to hold our data
	)
	res.Header().Add("Content-type", "application/json")
	if err := jsonDecoder.Decode(message); err != nil { //decode our data into our struct
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	pointless := BusinessLogic(message.Number)
	if err := jsonEncoder.Encode(pointless); err != nil { //encode our data and write it back to the response stream
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//Setup our simple router
func router() http.Handler {
	http.HandleFunc("/api/echo", Echo)
	return http.DefaultServeMux //this is a stdlib http.Handler
}

func main() {
	router := router()
	//start our server on port 3001
	if err := http.ListenAndServe(":3001", router); err != nil {
		log.Fatal(err)
	}
}
