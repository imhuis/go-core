package interface_method

import "fmt"

type Integer int

type Math interface {
	Add(i Integer)
	Multiply(i Integer) Integer
}

func (i2 *Integer) Add(i Integer) {
	*i2 = *i2 + i
}

func (i2 *Integer) Multiply(i Integer) Integer {
	return *i2 * i
}

func print() {
	var a Integer = 1
	var m Math = &a
	m.Add(1)
	fmt.Printf("1 + 2 = %d\n", m)
}
