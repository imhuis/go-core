package channel

import (
	"fmt"
	"time"
)

func chanClose() {
	c := make(chan int, 3)

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
