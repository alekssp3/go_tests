package main

import "fmt"

func main() {
	a := make(map[int]int)
	fmt.Print(a)
	a[1] = 1
	fmt.Print(a)
}