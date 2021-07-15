package main

import "fmt"

func main(){
	// sl := make([]int, 10)
	sl := make([]int, 3, 10)
	fmt.Println(sl, len(sl), cap(sl))

	sl = append(sl, 1)
	fmt.Println(sl, len(sl), cap(sl))

	sl = append(sl, []int{2, 3, 4, 5}...)
	fmt.Println(sl, len(sl), cap(sl))

	m1 := map[string]string{}
	m1["key"] = "value"
	fmt.Println(m1)
}