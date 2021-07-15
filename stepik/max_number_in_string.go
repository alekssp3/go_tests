package main

import . "fmt"

func main() {
	var s string
	var max = rune(0)
	Scan(&s)
	for _, c := range s {
		if c > max {
			max = c
		}
	}
	Print(string(max))
}