package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	c <- 1
	c <- 2

	close(c)
	time.Sleep(2 * time.Second)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
