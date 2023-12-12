package main

import (
	"fmt"
	"go-core/functions"
)

func main() {
	f := functions.Squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
