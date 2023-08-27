package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
)

const (
	name = "menglu"
	c    = iota
	d    = iota
)

type student struct {
	Age int
}

func main() {
	fmt.Println([...]string{"123"} == [...]string{"123"})
	// 切片不可以比较
	//fmt.Println([]string{"1"} == []string{"1"})

}
