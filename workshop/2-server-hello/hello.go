package main

// To run:
// 	go run main.go
//	curl -XPOST http://localhost:3001/api/echo -H 'Content-type:application/json' -d '{"message":"Heisann", "name":"Hallgeir"}'


import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

//Message wraps a message and stamps it
type Message struct {
	Message string `json:"message"` //tells the decoder what to decode into
	Stamp   int64  `json:"stamp,omitempty"`
	Name string `json:"name"`
}

// start3 OMIT
//BusinessLogic does awesome BusinessLogic
func BusinessLogic(greeting string, name string) *Message {
	mess := &Message{}
	mess.Message = greeting + " " + name
	mess.Stamp = time.Now().Unix()
	return mess
}
// end3 OMIT

// start2 OMIT
//MakeGreeting to welcome a user
func MakeGreeting(res http.ResponseWriter, req *http.Request) {
	var (
		jsonDecoder = json.NewDecoder(req.Body) //decoder reading from the post body
		jsonEncoder = json.NewEncoder(res)      //encoder writing to the response stream
		message     = &Message{}                // something to hold our data
	)
	res.Header().Add("Content-type", "application/json")
	if err := jsonDecoder.Decode(message); err != nil { //decode our data into our struct
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	aGreeting := BusinessLogic(message.Message, message.Name)
	if err := jsonEncoder.Encode(aGreeting); err != nil { //encode our data and write it back to the response stream
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}
// end2 OMIT

// start1 OMIT
//Setup our simple router
func router() http.Handler {
	http.HandleFunc("/api/echo", MakeGreeting)
	return http.DefaultServeMux //this is a stdlib http.Handler
}

func main() {
	router := router()
	//start our server on port 3001
	if err := http.ListenAndServe(":3001", router); err != nil {
		log.Fatal(err)
	}
}
// end1 OMIT
