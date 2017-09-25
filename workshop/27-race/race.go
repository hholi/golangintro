package main

import (
	"fmt"
)

// startrace OMIT
func sum(ret *int, done chan int) {
	for i := 0; i < 1000000; i++ {
		*ret++
	}
	done <- 1
}

func main() {
	sumTotal := 0
	cDone := make(chan int)
	go sum(&sumTotal, cDone)
	<-cDone
	fmt.Println("Total: ", sumTotal)
}

// endrace OMIT
