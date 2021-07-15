package main

import "fmt"

func f1() <- chan struct{} {
	done := make(chan struct{})
	go func() {
		fmt.Println("hello")
		close(done)
	}()
	return done
}

func main() {
	<- f1()
}
