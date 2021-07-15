package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Goroutine %d is started\n", id)
	dur := time.Duration(rand.Intn(3))
	time.Sleep(dur * time.Second)
	fmt.Printf("Goroutine %d stopped. Dur: %d\n", id, dur)
}

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(i, wg)
	}
	wg.Wait()
	fmt.Println("Done")
}
