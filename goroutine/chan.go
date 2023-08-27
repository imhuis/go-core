package goroutine

import "fmt"

func printChan() {
	// chan int
	var intChan chan int = make(chan int, 3)
	fmt.Printf("%T\n %v", intChan, intChan)
}
