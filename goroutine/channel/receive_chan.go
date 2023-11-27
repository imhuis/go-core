package channel

import "fmt"

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func receive(ch <-chan int) {
	fmt.Printf("receive%v", <-ch)
}
