package slice

import (
	"fmt"
	"testing"
)

func Test_Demo2(t *testing.T) {
	Demo2()
}

func Test_String(t *testing.T) {
	// len()返回的是byte数
	// rune是int32的别名
	var s string = "Hello,🏝"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}
