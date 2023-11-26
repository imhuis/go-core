package main

func hello() {
	println("Hello world goroutine")
}

func helloChan(done chan bool) {
	println("Hello world goroutine")
	done <- true
}

// 所有goroutine都会在main goroutine结束时结束
func main() {
	done := make(chan bool)
	//go hello()
	go helloChan(done)
	println("In main goroutine")
	<-done
	println("Main goroutine has finished")
}
