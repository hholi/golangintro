// start1 OMIT
package main

import "math"
import "fmt"

type Shape interface {
	Area() float64
}

type Square struct { // Note: no "implements" declaration
	side float64
}

func (sq Square) Area() float64 { return sq.side * sq.side }

// cont.
// end1 OMIT
// start2 OMIT
type Circle struct { // No "implements" declaration here either
	radius float64
}

func (c Circle) Area() float64 { return math.Pi * math.Pow(c.radius, 2) }

func main() {
	aCircle := Circle{1.0}
	aSquare := Square{2.0}
	shapes := []Shape{aCircle, aSquare}
	var sum float64
	for _, shape := range shapes {
		sum += shape.Area()
	}
	fmt.Println("Area of shapes is ", sum)
}
// end2 OMIT