package method

import "fmt"

type A struct {
}

type B struct {
}

// 类型断言
func typeAssertion() {
	var a interface{} = A{}
	var b interface{} = B{}
	if aa, ok := a.(A); ok {
		fmt.Printf("%v\t%v\n", aa, ok)
	}
	if bb, ok := b.(A); ok {
		fmt.Printf("%v\t%v", bb, ok)
	}
}
