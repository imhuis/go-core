package object_oriented

import "fmt"

/**
组合示例
*/

type People struct {
}

func (p *People) ShowA() {
	fmt.Println("student showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("student showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func mainMethod() {
	t := Teacher{}
	t.ShowA() // 调用student的个方法
	t.ShowB()
}
