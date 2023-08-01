package main

import (
	"fmt"
	_ "fmt"
)

func main1() {
	//x := 1
	//p := &x
	//fmt.Printf("x's pointer address %v\n", p)
	//*p = 2
	//fmt.Printf("x's value is %v\n", x)
	//fmt.Printf("p's value is %v\n", *p)

	var p1 = f()
	var p2 = f()
	fmt.Printf("p1's pointer address is %v\n", p1)
	fmt.Printf("p2's pointer address is %v\n", p2)

	fmt.Printf("f() == f() %v\n", p1 == p2)

	v := 1
	incr(&v)
	incr(&v)
	fmt.Printf("v = %v\n", v)
}

func f() *int {
	v := 1
	return &v
}

func incr(v *int) int {
	*v++
	return *v
}
