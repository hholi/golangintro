package main

import "fmt"

// startrace OMIT
func sum(sumChan chan int, done chan int) {
	for i := 0; i < 1000000; i++ {
		cur := <-sumChan
		cur++
		sumChan <- cur
	}
	done <- 1
}

func main() {
	cSum := make(chan int, 1) // Observe, buffered channel, place for 1
	cDone := make(chan int)
	cSum <- 0
	go sum(cSum, cDone)
	go sum(cSum, cDone)
	go sum(cSum, cDone)
	go sum(cSum, cDone)
	finishedTasks := <-cDone + <-cDone + <-cDone + <-cDone
	total := <-cSum
	fmt.Println("Total: ", total)
	fmt.Println("Tasks completed: ", finishedTasks)
}

// endrace OMIT
