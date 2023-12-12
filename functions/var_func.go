package functions

import (
	"fmt"
)

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func varFunc() {
	f := square
	fmt.Println(f(1))

	//f = add	error: 不可以赋值

	//var ff func(int) int
	//fmt.Println(ff((1)))	error:空指针不能运行

}

// 函数也是值

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}
