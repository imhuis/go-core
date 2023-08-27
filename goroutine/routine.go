package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type MyMutex struct {
	count int
	sync.Mutex
}

// #goroutines: 2
func main() {
	var ch chan int
	go func() {
		ch = make(chan int, 1)
		ch <- 1
	}()
	go func(ch chan int) {
		time.Sleep(time.Second)
		<-ch
	}(ch)
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
