package main

import . "fmt"

// 1,1,2,3,5

func main() {
	defer func(){recover()}()
	var n, tmp int
	var index = 1
	Scan(&n)
	if n == 1 || n == 2 {
		Print("1")
		panic("")
	}
	var prev, cur = 1, 1
	for index < n {
		tmp = prev
		prev = cur
		cur += tmp
		index++
	}
	Print(cur)
}
