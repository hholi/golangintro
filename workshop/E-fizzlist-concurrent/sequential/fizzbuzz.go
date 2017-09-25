package main

// To run:
// 	go run server.go
//	curl -XPOST http://localhost:3001/api/echo -H 'Content-type:application/json' -d '{"number":"15"}'

import (
	"fmt"
	"time"
)

type FizzBuzzResult struct {
	Number int 		`json:"number"` //tells the decoder what to decode into
	Result string 	`json:"result"` //tells the decoder what to decode into	
} 

//FizzBuzz transform number into text in mysterious ways
func FizzBuzz(i int) string {
	result := ""
	if i%3 == 0 {
		// Multiple of 3
		result = fmt.Sprintf("fizz")
	}
	if i%5 == 0 {
		// Multiple of 5
		result = fmt.Sprintf("%sbuzz", result)
	}

	if i%3 != 0 && i%5 != 0 {
		// Neither, so print the number itself
		result = fmt.Sprintf("%d", i)
	}
	time.Sleep(2*time.Second)
	return result
}

func FizzBuzzFromTo(from int, to int) []FizzBuzzResult {
	var allresults []FizzBuzzResult
	for number := from; number <= to; number++ {
		var result = FizzBuzzResult{number, FizzBuzz(number)}
		allresults = append(allresults, result)
	}
	return allresults
}
