// _Channels_ are the pipes that connect concurrent			OMIT
// goroutines. You can send values into channels from one	OMIT
// goroutine and receive those values into another			OMIT
// goroutine.												OMIT

package main

import "fmt"

func main() {

	// Create a new channel with `make(chan val-type)`.
	// Channels are typed by the values they convey.		OMIT
	messages := make(chan string)

	// _Send_ a value into a channel using the `channel <-` syntax
	// syntax. Here we send `"ping"`  to the `messages`		OMIT
	// channel we made above, from a new goroutine.			OMIT
	go func() { messages <- "ping" }()

	// The `<-channel` syntax _receives_ a value from the channel
	// channel. Here we'll receive the `"ping"` message		OMIT
	// we sent above and print it out.						OMIT
	msg := <-messages
	fmt.Println(msg)
}
