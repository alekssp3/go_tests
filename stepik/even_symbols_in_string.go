package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	var s = []rune(a)
	for i, c := range s {
		if i % 2 != 0 {
			fmt.Print(string(c))
		}
	}
}