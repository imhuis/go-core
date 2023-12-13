package main

import "time"

func main() {
	go func() {

	}()

	time.AfterFunc(5*time.Second, func() {
		println("after func")
	})

	select {
	case <-time.After(10 * time.Second):
	}

}
