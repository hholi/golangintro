package main

import (
	"fmt"
)

const fizz string = "Fizz"
const buzz string = "Buzz"

func fizzbuzz(cnt chan int, msg chan string) {
	for {
		i := <-cnt
		switch {
		case i%15 == 0:
			msg <- fmt.Sprintf("%d %s", i, fizz+buzz)
		case i%3 == 0:
			msg <- fmt.Sprintf("%d %s", i, fizz)
		case i%5 == 0:
			msg <- fmt.Sprintf("%d %s", i, buzz)
		default:
			msg <- fmt.Sprintf("%d", i)
		}
	}
}

func main() {
	count := make(chan int)
	message := make(chan string)
	go fizzbuzz(count, message)
	for i := 1; i < 101; i++ {
		count <- i
		s := <-message
		fmt.Println(s)
	}

}
