package funcs

import "fmt"

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func add(x, y int) int {
	return x + y
}

func varFunc() {
	f := square
	fmt.Println(f(1))

	//f = add	error: 不可以赋值

	//var ff func(int) int
	//fmt.Println(ff((1)))	error:空指针不能运行

}
