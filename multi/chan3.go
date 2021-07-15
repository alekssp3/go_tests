package main

import "fmt"

func f2(done chan struct{}) {
	fmt.Println("hello")
	close(done)
}

func main() {
	done := make(chan struct{})
	go f2(done)
	<- done
}
