package main

import (
	"fmt"
	"sync"
)

//func work(id int, wg *sync.WaitGroup, f func(id int), arg int) {
//	defer wg.Done()
//	fmt.Printf("Goroutine %d is started\n", id)
//	dur := time.Duration(rand.Intn(3))
//	time.Sleep(dur * time.Second)
//	f(arg)
//	fmt.Printf("Goroutine %d stopped. Dur: %d\n", id, dur)
//}

func f6(id int) {
	fmt.Printf("%d - hello\n", id)
}

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func (i int, f func(id int), wg *sync.WaitGroup) {
			f(i)
			defer wg.Done()
		}(i, f6, wg)
	}
	wg.Wait()
	fmt.Println("Done")
}
