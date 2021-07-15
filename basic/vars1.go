package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	a := "string"
	fmt.Printf("%T - %v", a, a)
	ioutil.ReadDir(".")
	ioutil.ReadDir(a)
	fmt.Printf("%s\n", a)
	test()
}

// test function
func test()  {
	fmt.Printf("Test")
}
