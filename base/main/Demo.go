package main

import "fmt"

func main1() {
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)
	var p = struct {
	}{}
	var q = struct {
	}{}
	fmt.Println(&p == &q)
}

func delta(old, new int) int {
	x := new(int)
	return new - x
}
