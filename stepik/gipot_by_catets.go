package main

import . "fmt"
import . "math"

func main() {
	var a, b float64
	Scan(&a, &b)
	Print(Pow(a * a + b * b, 0.5))
}