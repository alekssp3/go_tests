package main

import "fmt"

func main() {
	defer SaveData()
	fmt.Println("work")
	test(0)
}

func test(i int)  {
	if i == 0 {
		panic("i = 0")
	}
}

func SaveData()  {
	fmt.Println("save")
}