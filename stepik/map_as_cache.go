package main

import "fmt"

func main() {
	var a [10]int
	var m = make(map[int]int)
	for i := 0; i < 10; i++ {
		fmt.Scan(&a[i])
	}

	for _, i := range a {
		if _, ok := m[i]; ok {
		} else {
			m[i] = work(i)
		}
		fmt.Print(m[i], " ")
	}
}

func work(x int) (out int) {
	if x < 4 {
		out = x - 1
	} else {
		out = x + 1
	}
	return
}