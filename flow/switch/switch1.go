package _switch

import "fmt"

func main2() {
	switch num := 15; {
	case num < 50 && num < 51:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200", num)
	}
}
