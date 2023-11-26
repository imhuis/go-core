package channel

import "fmt"

func printChan() {
	// chan int
	var intChan chan int = make(chan int, 3)
	fmt.Printf("%T\n %v", intChan, intChan)
}

func sendData(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	close(ch)
}
