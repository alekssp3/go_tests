package main

import (
	"fmt"
	"math"
)

var k float64 = 1296
var p float64 = 6
var v float64 = 6

func T() float64 {
	return 6 / W()
}


func W() float64 {
	return math.Sqrt(k / M())
}

func M() float64 {
	return p * v
}

func main() {
	fmt.Print(T())
}
