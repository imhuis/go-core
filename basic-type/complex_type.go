package basictype

import "fmt"

func complexDemo() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x * y)
	// 实部
	fmt.Println(real(x * y))
	// 虚部
	fmt.Println(imag(x * y))
	// -1+0i
	fmt.Println(1i * 1i)
}
