package main

import "fmt"

func f5() {
	fmt.Println("hello")
}

func main() {
	done := make(chan struct{})
	go func(c chan struct{}) {
		f5()
		close(c)
	}(done)
	<- done
}
