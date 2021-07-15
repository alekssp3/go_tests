package main

import . "fmt"

// 1,1,2,3,5

func main() {
	var n, tmp int
	var index = 1
	Scan(&n)
	var prev, cur = 1, 1
	for cur < n {
		tmp = prev
		prev = cur
		cur += tmp
		index++
		// Printf("%d %d\n", cur, index)
	}
	if n == cur {
		Printf("%d\n", index)
	} else {
		Print("NO")
	}
}
