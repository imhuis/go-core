package interface_method

import "fmt"

func Example_NumbersAdd() {
	var a Integer = 1
	var m Math = &a
	m.Add(1)
	fmt.Printf("1 + 2 = %d\n", m)
	// output:
	// 1 + 1 = 2
}
