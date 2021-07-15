package main

import (
	"fmt"
)


func main()  {
	fmt.Print(minimumFromFour())
}

func minimumFromFour() int {
	var cur, min int
	fmt.Scan(&min)
	// var ar = []int{0, 0, 0, 0}
	for i := 0; i < 3; i++ {
		fmt.Scan(&cur)
		if cur < min {
			min = cur
		}
	}
	return min
}
