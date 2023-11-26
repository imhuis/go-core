package goroutine

import (
	"fmt"
	"time"
)

func Counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		time.Sleep(1 * time.Second)
		out <- x
	}
	close(out)
}

func Squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func Printer(in <-chan int) {
	for v := range in {
		fmt.Print(v)
	}
}

// count a fib number
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
