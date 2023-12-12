package interface_method

import (
	"fmt"
	"testing"
)

func Test_InterfaceMethod(t *testing.T) {
	addr := Addr{start: 1}
	addr.AddTo(2)
	if addr.start != 3 {
		t.Errorf("addr.start should be 3, but get %d", addr.start)
	}
	addr.AddTo(3)
	if addr.start != 6 {
		t.Errorf("addr.start should be 6, but get %d", addr.start)
	}

	// 方法即函数
	ff := addr.AddTo
	fmt.Println(ff(1))
}
