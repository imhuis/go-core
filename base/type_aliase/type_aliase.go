package type_aliase

import "fmt"

type I int
type D = int

func test() {
	var v interface{}
	switch i := v.(type) {
	case int:
		fmt.Printf("", i)
		//case D: 重复

	}
}

type T1 struct {
}

type T2 = T1

func (t *T1) say() {

}

func (t *T2) hell0() {

}

func test2() {
	var t1 T1
	var t2 T2

	t1.say()
	t1.hell0()
	t2.say()
	t2.hell0()
}
