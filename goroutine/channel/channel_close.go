package channel

import (
	"fmt"
	"time"
)

func chanCloseOK() {
	c := make(chan int, 2)
	c <- 1
	v, ok := <-c
	time.Sleep(2 * time.Second)
	fmt.Println(v, ok)
	close(c)
	v, ok = <-c
	fmt.Println(v, ok)
}

func chanClose() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	close(c)

	time.Sleep(10 * time.Second)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
