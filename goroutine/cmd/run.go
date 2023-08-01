package main

import "fmt"

func main() {

	var intchan chan int = make(chan int, 3)
	fmt.Printf("%T\n %v", intchan, intchan)
}
