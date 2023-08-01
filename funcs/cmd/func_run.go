package main

import (
	"fmt"
	"learning-go/funcs"
)

func main() {
	f := funcs.Squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
