package slice

import "fmt"

func compare() {
	fmt.Println([...]string{"123"} == [...]string{"123"})
	// 切片不可以被比较
	//fmt.Println([]string{"1"} == []string{"1"})
}
