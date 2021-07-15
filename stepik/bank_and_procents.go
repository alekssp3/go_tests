package main

import "fmt"


func main()  {
	var x, p, y, c int
	fmt.Scan(&x, &p, &y)
	for x <= y {
		x += (x * p / 100)
		c++
	}
	fmt.Printf("%d\n", c)
	
}
