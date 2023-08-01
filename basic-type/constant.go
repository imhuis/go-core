package basictype

import (
	"fmt"
	"math"
)

type Weekday int

const (
	// 除第一项外，其他项等号右边表达式可以省略，以为复用前面一项
	A = 1
	B
	C = 2
	D

	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func demo1() {
	pi := math.Pi
	fmt.Print(pi)
}
