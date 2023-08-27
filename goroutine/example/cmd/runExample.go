package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(number int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(number)
		}(i, c)
	}

	if WaitTimeout(&wg, 5*time.Second) {
		close(c)
		fmt.Printf("timeout exit")
	}
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	return false
}
