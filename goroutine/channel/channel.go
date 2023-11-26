package channel

import "fmt"

func printChanMeta() {
	var intChan chan int = make(chan int, 3)
	fmt.Printf("%T\n", intChan)
}

func printChan() {
	// chan int
	var intChan chan int = make(chan int, 2)
	// 双向通道可以转换为单向通道
	sendOnlyChan(intChan)
}

func sendData(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	close(ch)
}

func sendOnlyChan(ch chan<- int) {
	ch <- 1
	ch <- 2
	close(ch)
	//<- ch // invalid operation: <-ch (receive from send-only type chan<- int)
}
