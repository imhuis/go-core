package main

import "fmt"

// 数组

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	//f1()
	f2()
}

var array = [1]int{}

func f1() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])
}

func f2() {
	r := [...]int{99: -1}
	fmt.Printf("%T", r)
}

func f3() {
	array2 := [2]int{1, 2}
	fmt.Println(array2)
	array3 := [...]int{1, 2, 4}
	fmt.Println(array3)
}
