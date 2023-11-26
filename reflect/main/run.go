package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Printf("%T\n", t)
	fmt.Println(t.Size(), t.String())
}
