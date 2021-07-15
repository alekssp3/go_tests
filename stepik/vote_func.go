package main

import "fmt"


func main()  {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Print(vote(a, b, c))
}

// vote ...
func vote(a, b, c int) int {
	var out int
	if a == b || a == c {
		out = a
	}

	if b == c {
		out = b
	}
	return out
}
