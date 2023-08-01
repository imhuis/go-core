package funcs

func Squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
