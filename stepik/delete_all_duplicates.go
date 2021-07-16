package main

import (
	"fmt"
	"strings"
)


func main() {
	var a string
	fmt.Scan(&a)
	var s = []rune(a)
	for _, c := range s {
		if strings.Count(a, string(c)) == 1 {
			fmt.Print(string(c))
		}
	}
}