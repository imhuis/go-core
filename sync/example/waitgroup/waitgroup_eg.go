package waitgroup

import (
	"sync"
	"time"
)

// panic: sync: WaitGroup is reused before previous Wait has returned
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}
