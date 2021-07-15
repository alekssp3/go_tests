package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pinger(c chan string) {
	for {
		c <- "ping"
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	}
}

func ponger(c chan string) {
	for {
		c <- "pong"
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	c := make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var s string
	fmt.Scan(&s)
}
