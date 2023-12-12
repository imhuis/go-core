package main

import (
	"go-core/method"
)

/*
	测试方法的调用方式
*/

func main() {
	p := method.Point{
		X: 1,
		Y: 2,
	}
	// 变量 method.Point
	p.ScaleBy(4)

	// 指针类型 *method.Point
	pptr := &p
	pptr.ScaleBy(5)
	(&p).ScaleBy(5)

	// compile error 不能获取类型字面量的地址
	//method.Point{1, 2}.ScaleBy(2)
	//
	method.Point{1, 2}.Distance(p)

}
