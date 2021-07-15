package main

import (
	"fmt"
	"time"
)

//func main() {
//	t := time.NewTimer(time.Second)
//	go func() {
//		<- t.C
//	}()
//	t.Stop()
//	t.Reset(time.Second * 2)
//	<- t.C
//}

func main() {
	<- work()
}

func work() <- chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		stop := time.NewTimer(time.Second)
		tick := time.NewTicker(time.Millisecond * 200)
		defer tick.Stop()
		for {
			select {
			case <-stop.C:
				return
			case <-tick.C:
				fmt.Println("tik-tok")
			}
		}
	}()
	return done
}