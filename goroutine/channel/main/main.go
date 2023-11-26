package main

func main() {
	sendch := make(chan int)
	go send(sendch)

	<-sendch
}

func send(ch chan<- int) {
	ch <- 1
}
