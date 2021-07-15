package main

import . "fmt"

func main() {
	v := 5
	p := &v
	Print(*p, " ")
	changePointer(p)
	Print(*p, " ")
}

func changePointer(p *int) {
	v := 3
	p = &v
}