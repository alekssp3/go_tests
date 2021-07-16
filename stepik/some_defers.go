package main

import "fmt"

func main() {
	var a int = 5
	defer foo(a)
	a = 6
	fmt.Print(a)
}

func foo(i int) int {
	fmt.Print(i)
	i++
	return i
}