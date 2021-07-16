package main

import 	"fmt"

func main() {
	fmt.Print(sumInts(1, 2, 3))
	fmt.Print(sumInts(1, 2))
}

func sumInts(a ...interface{}) (s, c int)  {
	var sum, cnt int
	for _, i := range a {
		sum += i.(int)
		cnt++
	}
	return sum, cnt
}