package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go f(i)
	}
	var s string
	fmt.Scan(&s)
}

func f(n int) {
	for i := 0; i < 3; i++ {
		fmt.Println(n, " : ", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}