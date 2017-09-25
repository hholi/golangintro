package main

import "fmt"

func divMod(den, div int) (quo, rem int) {
	return den / div, den % div
}

func main() {
	quo, rem := divMod(21, 4)
	fmt.Println("Quotient:  ", quo)
	fmt.Println("Remainder: ", rem)
}
