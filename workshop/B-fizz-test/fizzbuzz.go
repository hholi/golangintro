package main

import "fmt"

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
	return result
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
