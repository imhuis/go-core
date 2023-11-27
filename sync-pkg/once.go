package sync_pkg

import (
	"fmt"
	"sync"
)

func Once() error {
	var once sync.Once
	onceFunc := func() {
		fmt.Printf("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			// 只会执行一遍
			once.Do(onceFunc)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
	return nil
}
