package goroutine

import (
	"fmt"
	"time"
)

func demo1() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			time.Sleep(2 * time.Second)
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

// 拆分
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

func demo2() {
	//naturals := make(chan int)
	//squares := make(chan int)
	//go goroutine.Counter(naturals)
	//go goroutine.Squarer(squares, naturals)
	//goroutine.Printer(squares)
}
