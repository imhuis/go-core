package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func Test_Squarer(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)
	go Counter(naturals)
	go Squarer(squares, naturals)
	Printer(squares)
}

func Test_PrintSquarer(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			time.Sleep(1 * time.Second)
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	for {
		println(<-squares)
	}
}

func Test_fib(t *testing.T) {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	// Output:
	// Fibonacci(45) = 1134903170
}
