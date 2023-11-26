package channel

func Example_ChanClose() {
	chanClose()
	// Output:
	// 1
	// 2
	// 0
	// 0
	// 0
}

func Example_ChanCloseOK() {
	chanCloseOK()
	// Output:
	// 1 true
	// 0 false
}
