package main

import "fmt"

func main() {

}

func printChan() {
	// chan int
	var intChan chan int = make(chan int, 3)
	fmt.Printf("%T\n %v", intChan, intChan)
}
