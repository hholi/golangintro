package main

import (
	"log"
	"testing"
)

func TestFizzbuzz_1(t *testing.T) {
	result := FizzBuzz(1)
	if "1" != result {
		t.Fail()
		log.Println("expected the message to equal test")
	}
}

func TestFizzbuzz_3(t *testing.T) {
	result := FizzBuzz(3)
	if "fizz" != result {
		t.Fail()
		log.Println("expected the message to equal test")
	}
}

func TestFizzbuzz_5(t *testing.T) {
	result := FizzBuzz(5)
	if "buzz" != result {
		t.Fail()
		log.Println("expected the message to equal test")
	}
}

func TestFizzbuzz_15(t *testing.T) {
	result := FizzBuzz(15)
	if "fizzbuzz" != result {
		t.Fail()
		log.Println("expected the message to equal test")
	}
}
