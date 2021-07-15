package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()

	wg := new(sync.WaitGroup)
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, tick.C, wg)
	}
	wg.Wait()
}

func worker(i int, c <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	<-c
	fmt.Printf("worker %d ends self work\n", i)
}
