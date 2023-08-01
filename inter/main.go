package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 10*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("sleep for %v seconds", *period)
	time.Sleep(*period)
	fmt.Println()
}
