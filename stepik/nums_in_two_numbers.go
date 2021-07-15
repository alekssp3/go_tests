package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	for _, ac := range a {
		for _, bc := range b {
			if ac == bc {
				fmt.Print(string(ac), " ")
			}
		}
	}
}
