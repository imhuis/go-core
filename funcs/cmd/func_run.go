package main

import (
	"fmt"
	"go-core/funcs"
)

func main() {
	f := funcs.Squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
