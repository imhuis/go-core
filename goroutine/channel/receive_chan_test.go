package channel

import "fmt"

func Example_ReceiveChan() {
	ch := make(chan int)
	go produce(ch)

	for v := range ch {
		fmt.Printf("%v\t", v)
	}
	// Output:
	// 0	1	2	3	4	5	6	7	8	9
}
