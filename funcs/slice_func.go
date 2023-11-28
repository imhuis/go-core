package funcs

// 变长函数
func sum(vals ...int) {
	total := 0
	for _, val := range vals {
		total += val
	}
}

func sumTest() {
	// a is a slice
	a := []int{1, 2, 3}
	sum(a...)
	sum(1, 2, 3)
	sum()
}
