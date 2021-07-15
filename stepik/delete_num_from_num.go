package main

import . "fmt"

func main() {
	var s string
	var n string
	Scan(&s, &n)
	for _, c := range s {
		if string(c) != n {
			Print(string(c))
		}
	}
}