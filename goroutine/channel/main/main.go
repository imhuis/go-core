package main

import "fmt"

func main() {
	sendch := make(chan int)
	go send(sendch)

	rece := <-sendch
	fmt.Println(rece)
}

func send(ch chan<- int) {
	ch <- 1
}
