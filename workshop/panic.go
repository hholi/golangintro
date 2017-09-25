package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
		fmt.Println("Wrapping up")
	}()

	fmt.Println("(Doing something that may fail)")
	panic("PANIC!")
	fmt.Println("(Next steps that will never be executed)")
}
