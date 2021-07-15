package main

import . "fmt"

func main() {
	var s string
	Scan(&s)
	for _, c := range s {
		var d = c - rune('0')
		Print(d * d)
	}
}