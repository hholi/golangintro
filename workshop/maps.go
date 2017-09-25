package main

import (
	"fmt"
)

func main() {
	xArray := [5]float64{98, 93, 77, 82, 83}

	xSliceFirst := xArray[:len(xArray)/2]
	xSliceSecond := xArray[len(xArray)/2:]

	xSliceMap := make(map[string][]float64)
	xSliceMap["first"] = xSliceFirst
	xSliceMap["second"] = xSliceSecond

	fmt.Println(xSliceMap)
}
