package main

import (
	"fmt"
	"os"
)

func main() {
	//args := os.Args
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("%v\t", os.Args[i])
	}
}
